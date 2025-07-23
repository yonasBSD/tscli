// pkg/tscli/client.go
//
// Thin wrapper around tailscale-client-go that:
//
//   - picks up tailnet / api-key / debug from Viper
//   - logs every HTTP request & response when --debug or TSCLI_DEBUG=1 is set
package tscli

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"

	"github.com/jaxxstorm/tscli/pkg/version"
	"github.com/spf13/viper"
	tsapi "tailscale.com/client/tailscale/v2"
)

const (
	defaultBaseURL     = "https://api.tailscale.com"
	defaultContentType = "application/json"
)

// getUserAgent returns the properly formatted user agent string
func getUserAgent() string {
	return fmt.Sprintf("tscli/%s (Go client)", version.GetVersion())
}

func New() (*tsapi.Client, error) {
	tailnet := viper.GetString("tailnet")
	apiKey := viper.GetString("api-key")
	if tailnet == "" {
		return nil, fmt.Errorf("tailnet is required")
	}
	if apiKey == "" {
		return nil, fmt.Errorf("api-key is required")
	}

	userAgent := getUserAgent()

	// Create a custom transport that ensures UserAgent is always set
	transport := &userAgentTransport{
		rt:        http.DefaultTransport,
		userAgent: userAgent,
	}

	// Wrap with debug logging if enabled
	if viper.GetBool("debug") {
		transport.rt = &logTransport{rt: transport.rt}
	}

	httpClient := &http.Client{
		Transport: transport,
	}

	return &tsapi.Client{
		Tailnet:   tailnet,
		APIKey:    apiKey,
		UserAgent: userAgent,
		HTTP:      httpClient,
	}, nil
}

// Do performs an HTTP call on top of an existing *tsapi.Client.  Useful for
// endpoints not yet covered by the SDK.  When “debug” is on, full request /
// response dumps are printed to stderr.
func Do(
	ctx context.Context,
	c *tsapi.Client,
	method, path string,
	body any,
	out any,
) (http.Header, error) {

	base := c.BaseURL
	if base == nil {
		b, _ := url.Parse(defaultBaseURL)
		base = b
	}

	u, err := url.Parse(path)
	if err != nil {
		return nil, fmt.Errorf("invalid path: %w", err)
	}

	u.Path = strings.ReplaceAll(u.Path, "{tailnet}", url.PathEscape(c.Tailnet))

	full := base.ResolveReference(&url.URL{
		Path:     "/api/v2" + u.Path,
		RawQuery: u.RawQuery,
	})

	var rdr io.Reader
	if body != nil {
		switch v := body.(type) {
		case []byte:
			rdr = bytes.NewReader(v)
		case string:
			rdr = strings.NewReader(v)
		default:
			b, err := json.Marshal(v)
			if err != nil {
				return nil, fmt.Errorf("marshal body: %w", err)
			}
			rdr = bytes.NewReader(b)
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, full.String(), rdr)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("Accept", defaultContentType)
	if body != nil {
		req.Header.Set("Content-Type", defaultContentType)
	}
	if c.APIKey != "" {
		req.SetBasicAuth(c.APIKey, "")
	}

	// dump request information if debug is enabled
	if viper.GetBool("debug") {
		if dump, _ := httputil.DumpRequestOut(req, true); len(dump) > 0 {
			os.Stderr.Write(dump)
		}
	}

	httpc := c.HTTP
	if httpc == nil {
		httpc = http.DefaultClient
	}

	res, err := httpc.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return res.Header, err
	}

	// dump response information if debug is enabled
	if viper.GetBool("debug") {
		if dump, _ := httputil.DumpResponse(res, false); len(dump) > 0 {
			os.Stderr.Write(dump)
		}
		if len(data) < 4_096 {
			os.Stderr.Write(data)
			fmt.Fprintln(os.Stderr)
		}
	}

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusMultipleChoices {
		return res.Header, fmt.Errorf(
			"tailscale API %s %s -> %d: %s",
			method, path, res.StatusCode, strings.TrimSpace(string(data)),
		)
	}

	if out == nil || len(data) == 0 {
		return res.Header, nil
	}
	if err := json.Unmarshal(data, out); err != nil {
		return res.Header, fmt.Errorf("decode response: %w", err)
	}
	return res.Header, nil
}

type logTransport struct{ rt http.RoundTripper }

func (t *logTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if dump, _ := httputil.DumpRequestOut(req, true); len(dump) > 0 {
		os.Stderr.Write(dump)
	}
	resp, err := t.rt.RoundTrip(req)
	if err != nil {
		return resp, err
	}
	if dump, _ := httputil.DumpResponse(resp, false); len(dump) > 0 {
		os.Stderr.Write(dump)
	}
	return resp, nil
}

// userAgentTransport wraps an http.RoundTripper to ensure UserAgent is always set
type userAgentTransport struct {
	rt        http.RoundTripper
	userAgent string
}

func (t *userAgentTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Always set our custom user agent
	req.Header.Set("User-Agent", t.userAgent)
	return t.rt.RoundTrip(req)
}

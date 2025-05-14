// pkg/tscli/client.go

package tscli

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/spf13/viper"
	tsapi "tailscale.com/client/tailscale/v2"
)

const (
	defaultBaseURL     = "https://api.tailscale.com"
	defaultContentType = "application/json"
)

// New returns a fully configured *tsapi.Client using flags / env from viper.
func New() (*tsapi.Client, error) {
	tailnet := viper.GetString("tailnet")
	apiKey := viper.GetString("api-key")

	switch {
	case tailnet == "":
		return nil, fmt.Errorf("tailnet is required")
	case apiKey == "":
		return nil, fmt.Errorf("api-key is required")
	}

	return &tsapi.Client{
		Tailnet:   tailnet,
		APIKey:    apiKey,
		UserAgent: "tscli",
		HTTP:      &http.Client{},
	}, nil
}

// Do performs a raw HTTP call using an existing *tsapi.Client.
//
//   method – http.MethodGet / http.MethodPost / …
//   path   – part _after_ “…/api/v2”; "{tailnet}" will be replaced & escaped.
//   body   – JSON-serialisable value, []byte, string, or nil.
//   out    – pointer to decode JSON response into, or nil if no body expected.
//
// It returns the response headers so callers can inspect rate-limit info, etc.
func Do(
	ctx context.Context,
	c *tsapi.Client,
	method, path string,
	body any,
	out any,
) (http.Header, error) {

	// -------- build URL ------------------------------------------------------
	base := c.BaseURL
	if base == nil {
		var err error
		base, err = url.Parse(defaultBaseURL)
		if err != nil {
			return nil, fmt.Errorf("parse base URL: %w", err)
		}
	}

	path = strings.ReplaceAll(path, "{tailnet}", url.PathEscape(c.Tailnet))
	full := base.ResolveReference(&url.URL{Path: "/api/v2" + path})

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

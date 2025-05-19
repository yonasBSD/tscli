// cmd/tscli/get/device/cli_flags_test.go

package device

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"testing"

	tsapi "tailscale.com/client/tailscale/v2"
)


type dummyRT struct{ body []byte }

func (d *dummyRT) RoundTrip(*http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewReader(d.body)),
		Header:     make(http.Header),
	}, nil
}


func TestGetDeviceFlagValidation(t *testing.T) {
	t.Parallel()

	// Fake client returned when flags are correct.
	fakeDev := tsapi.Device{ID: "123", Hostname: "ok", OS: "linux"}
	b, _ := json.Marshal(fakeDev)
	stubClient := func() (*tsapi.Client, error) {
		base, _ := url.Parse("http://fake")
		return &tsapi.Client{
			BaseURL: base,
			HTTP:    &http.Client{Transport: &dummyRT{body: b}},
		}, nil
	}

	cases := []struct {
		name    string
		args    []string
		useStub bool
		wantErr bool
	}{
		{
			name:    "missing required flag",
			args:    []string{},
			wantErr: true,
		},
		{
			name:    "unknown flag",
			args:    []string{"--bogus"},
			wantErr: true,
		},
		{
			name:    "valid device flag",
			args:    []string{"--device", "123"},
			useStub: true, // must avoid real HTTP
			wantErr: false,
		},
	}

	for _, tc := range cases {
		tc := tc // capture
		t.Run(tc.name, func(t *testing.T) {
			orig := newClient
			if tc.useStub {
				newClient = stubClient
			}
			defer func() { newClient = orig }()

			cmd := Command()
			cmd.SetArgs(tc.args)
			cmd.SetOut(io.Discard)
			cmd.SetErr(io.Discard)

			err := cmd.ExecuteContext(context.Background())
			if tc.wantErr && err == nil {
				t.Fatalf("expected error, got nil")
			}
			if !tc.wantErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

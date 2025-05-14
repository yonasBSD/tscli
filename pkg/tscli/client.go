package tscli

import (
	"fmt"
	"net/http"

	"github.com/spf13/viper"
	tsapi "tailscale.com/client/tailscale/v2"
)

func New() (*tsapi.Client, error) {
	tailnet := viper.GetString("tailnet")
	apiKey := viper.GetString("api-key")
	if tailnet == "" {
		return nil, fmt.Errorf("tailnet is required")
	}
	if apiKey == "" {
		return nil, fmt.Errorf("api-key is required")
	}

	httpClient := &http.Client{}

	return &tsapi.Client{
		Tailnet:   tailnet,
		APIKey:    apiKey,
		UserAgent: "tscli",
		HTTP:      httpClient,
	}, nil
}

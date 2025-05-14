package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/jaxxstorm/tscli/cmd/tscli/devices"
	viper "github.com/spf13/viper"

	"github.com/jaxxstorm/tscli/pkg/contract"
)

var (
	api_key string
	tailnet string
)

func configureCLI() *cobra.Command {

	viper := viper.GetViper()

	rootCommand := &cobra.Command{
		Use:  "tscli",
		Long: "A cli tool for interacting with the Tailscale API.",
	}

	rootCommand.AddCommand(
		devices.Command())

	rootCommand.PersistentFlags().StringVarP(&api_key, "api-key", "k", "", "Tailscale API key.")
	rootCommand.PersistentFlags().StringVarP(&tailnet, "tailnet", "n", "-", "Tailscale tailnet.")
	rootCommand.MarkFlagRequired("api-key")
	rootCommand.MarkFlagRequired("tailnet")

	viper.AutomaticEnv()

	viper.BindEnv("api-key", "TAILSCALE_API_KEY")
	viper.BindEnv("tailnet", "TAILSCALE_TAILNET")
	viper.BindPFlag("api-key", rootCommand.PersistentFlags().Lookup("api-key"))
	viper.BindPFlag("tailnet", rootCommand.PersistentFlags().Lookup("tailnet"))


	return rootCommand
}

func main() {
	rootCommand := configureCLI()

	

	if err := rootCommand.Execute(); err != nil {
		contract.IgnoreIoError(fmt.Fprintf(os.Stderr, "%v\n", err))
		os.Exit(1)
	}
}
package main

import (
	"fmt"
	"os"

	"github.com/jaxxstorm/tscli/cmd/tscli/delete"
	"github.com/jaxxstorm/tscli/cmd/tscli/get"
	"github.com/jaxxstorm/tscli/cmd/tscli/list"
	"github.com/jaxxstorm/tscli/cmd/tscli/set"
	"github.com/spf13/cobra"
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
		get.Command())
	rootCommand.AddCommand(
		list.Command())
	rootCommand.AddCommand(
		delete.Command())
	rootCommand.AddCommand(
		set.Command())

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

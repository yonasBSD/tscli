// cmd/tscli/devices/posture/command.go

package posture

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/jaxxstorm/tscli/pkg/tscli"
	"github.com/spf13/cobra"
	tsapi "tailscale.com/client/tailscale/v2"
)

var keyRe = regexp.MustCompile(`^custom:[A-Za-z0-9_:]{1,43}$`) // 7 + 43 = 50
var strValRe = regexp.MustCompile(`^[A-Za-z0-9_.]{1,50}$`)

func coerceValue(in string) (any, error) {
	if b, err := strconv.ParseBool(in); err == nil {
		return b, nil
	}
	if i, err := strconv.ParseInt(in, 10, 64); err == nil {
		return i, nil
	}
	if !strValRe.MatchString(in) {
		return nil, errors.New(`string value may be 1-50 chars: letters, numbers, underscores, periods`)
	}
	return in, nil
}

func parseExpiry(exp string) (tsapi.Time, error) {
	if exp == "" {
		return tsapi.Time{}, nil
	}
	t, err := time.Parse(time.RFC3339, exp)
	if err != nil {
		return tsapi.Time{}, err
	}
	return tsapi.Time{Time: t.UTC()}, nil
}

func Command() *cobra.Command {

	var (
		keyFlag     string
		valFlag     string
		expiryFlag  string
		commentFlag string
		deviceFlag  string

		parsedValue any
		parsedExp   tsapi.Time
	)

	command := &cobra.Command{
		Use:   "posture",
		Short: "Set a custom posture attribute on a device",
		Long:  "Set (or update) a custom:<attr> posture attribute on the specified device.",

		// validate inputs for the command
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// key
			if !keyRe.MatchString(keyFlag) {
				return errors.New(`--key must start with "custom:" and contain only letters, numbers, underscores or colons (≤ 50 chars)`)
			}
			// value
			v, err := coerceValue(valFlag)
			if err != nil {
				return err
			}
			parsedValue = v

			// expiry
			exp, err := parseExpiry(expiryFlag)
			if err != nil {
				return fmt.Errorf("invalid --expiry: %w", err)
			}
			parsedExp = exp

			// device
			if deviceFlag == "" {
				return errors.New("--device is required")
			}
			return nil
		},

		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := tscli.New()
			if err != nil {
				return fmt.Errorf("failed to create client: %w", err)
			}

			req := tsapi.DevicePostureAttributeRequest{
				Value:   parsedValue,
				Expiry:  parsedExp,
				Comment: commentFlag,
			}

			if err := client.Devices().SetPostureAttribute(
				context.Background(),
				deviceFlag,
				keyFlag,
				req,
			); err != nil {
				return fmt.Errorf("failed to set posture attribute: %w", err)
			}

			payload := map[string]string{
				"result": fmt.Sprintf("device %s: %s set to %v", deviceFlag, keyFlag, parsedValue),
			}
			out, _ := json.MarshalIndent(payload, "", "  ")
			fmt.Fprintln(os.Stdout, string(out))
			return nil
		},
	}

	command.Flags().StringVar(&deviceFlag, "device", "", "Device ID to update")
	command.Flags().StringVar(&keyFlag, "key", "", `Posture attribute key (must start with "custom:")`)
	command.Flags().StringVar(&valFlag, "value", "", "Value to set (string, number, or boolean)")
	command.Flags().StringVar(&expiryFlag, "expiry", "", `Optional RFC-3339 expiry time (e.g. "2025-01-02T15:04:05Z")`)
	command.Flags().StringVar(&commentFlag, "comment", "", "Optional audit-log comment")

	_ = command.MarkFlagRequired("device")
	_ = command.MarkFlagRequired("key")
	_ = command.MarkFlagRequired("value")

	return command
}

// Package vpn manages the vpn command
package vpn

import (
	"log/slog"

	"ouroboros/cmd/vpn/start"
	"ouroboros/cmd/vpn/stop"

	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "vpn",
		Short: "vpn options",
		Args:  cobra.NoArgs,
	}

	cmd.AddCommand(stop.Cmd())
	cmd.AddCommand(start.Cmd())

	slog.Debug("command:", "value", cmd.Use)

	return &cmd
}

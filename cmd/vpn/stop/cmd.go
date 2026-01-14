// Package stop sets up the stop command
package stop

import (
	"log/slog"

	stop "ouroboros/internal/vpn/stop"

	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "stop",
		Short: "stop the vpn",
		RunE:  stop.Cmd,
		Args:  cobra.NoArgs,
	}

	slog.Debug("command stop:", "value", cmd.Use)

	return &cmd
}

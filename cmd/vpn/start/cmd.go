// Package start sets up the start command
package start

import (
	"log/slog"

	start "ouroboros/internal/vpn/start"

	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "start",
		Short: "start the vpn",
		RunE:  start.Cmd,
		Args:  cobra.NoArgs,
	}

	slog.Debug("command start:", "value", cmd.Use)

	return &cmd
}

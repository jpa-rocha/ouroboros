package stop

import (
	"context"
	"log/slog"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// Cmd is the command handler for stopping the VPN connection.
// Takes a Cobra command and ignores the arguments.
// Calls Action with the command's context and returns its result.
func Cmd(cmd *cobra.Command, _ []string) error {
	return Action(cmd.Context())
}

// Action disconnects the VPN using vpnc-disconnect.
// Takes context for logging.
// Logs any errors as warnings and always returns nil (graceful handling of disconnection).
// Logs final status indicating the VPN is not connected.
func Action(ctx context.Context) error {
	cmdStop := exec.Command("sudo", "vpnc-disconnect")

	out, err := cmdStop.CombinedOutput()
	if err != nil {
		slog.WarnContext(ctx, "ouroboros", "error", strings.TrimRight(string(out), "\n"))
	}

	slog.InfoContext(ctx, "ouroboros", "status", "not connected")

	return nil
}

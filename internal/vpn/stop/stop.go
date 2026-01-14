package stop

import (
	"context"
	"log/slog"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

func Cmd(cmd *cobra.Command, _ []string) error {
	return Action(cmd.Context())
}

func Action(ctx context.Context) error {
	cmdStop := exec.Command("sudo", "vpnc-disconnect")

	out, err := cmdStop.CombinedOutput()
	if err != nil {
		slog.WarnContext(ctx, "ouroboros", "error", strings.TrimRight(string(out), "\n"))
	}

	slog.InfoContext(ctx, "ouroboros", "status", "not connected")

	return nil
}

// Package start holds the internal of the start command
package start

import (
	"context"
	"log/slog"
	"os/exec"
	"strconv"
	"time"

	config "ouroboros/internal/config"
	stop "ouroboros/internal/vpn/stop"

	"github.com/spf13/cobra"
)

func Cmd(cmd *cobra.Command, _ []string) error {
	ctx := cmd.Context()

	_ = stop.Action(ctx)
	time.Sleep(time.Duration(time.Second * 1))

	return Start(ctx)
}

func Start(ctx context.Context) error {
	cmdStart := exec.Command(
		"sudo",
		"vpnc",
		"--gateway",
		config.Opt.Settings.Address,
		"--id",
		config.Opt.Secrets.VPN.IPSec.Username,
		"--secret",
		config.Opt.Secrets.VPN.IPSec.Password,
		"--username",
		config.Opt.Secrets.VPN.ID.Username,
		"--password",
		config.Opt.Secrets.VPN.ID.Password,
		"--ifmtu",
		strconv.Itoa(config.Opt.Settings.MTU),
	)

	slog.InfoContext(ctx, "ouroboros", "status", "starting")

	out, err := cmdStart.CombinedOutput()
	if err != nil {
		slog.ErrorContext(ctx, "ouroboros", "error", string(out))
		slog.ErrorContext(ctx, "ouroboros:", "error", err.Error())

		return err
	}

	slog.InfoContext(ctx, "ouroboros", "status", out)

	slog.InfoContext(ctx, "ouroboros", "status", "connected")

	return nil
}

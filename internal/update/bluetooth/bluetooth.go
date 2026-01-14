// Package bluetooth has the internal for the bluetooth command
package bluetooth

import (
	config "ouroboros/internal/config"
	update "ouroboros/internal/update"

	"github.com/spf13/cobra"
)

func RunBluetooth(cmd *cobra.Command, _ []string) error {
	ctx := cmd.Context()
	update.InstallPrereqs(ctx, false)
	update.InstallDriver(
		ctx,
		update.BLUETOOTH,
		update.INSTALL_BLUETOOTH,
		config.Opt.Repos.BlueToothGitRepo,
	)

	if config.Opt.Reboot.Yes {
		if err := update.RebootCmd(ctx); err != nil {
			return err
		}
	}

	return nil
}

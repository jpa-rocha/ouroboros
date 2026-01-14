// Package all holds the internals of the all command
package all

import (
	config "ouroboros/internal/config"
	update "ouroboros/internal/update"
	"ouroboros/internal/update/camera"

	"github.com/spf13/cobra"
)

func RunAll(cmd *cobra.Command, _ []string) error {
	ctx := cmd.Context()
	update.InstallPrereqs(ctx, true)

	update.InstallDriver(
		ctx,
		update.AUDIO,
		update.INSTALL_AUDIO,
		config.Opt.Repos.AudioGitRepo,
	)

	update.InstallDriver(
		ctx,
		update.BLUETOOTH,
		update.INSTALL_BLUETOOTH,
		config.Opt.Repos.BlueToothGitRepo,
	)

	camera.InstallCamera(ctx)

	if config.Opt.Reboot.Yes {
		if err := update.RebootCmd(ctx); err != nil {
			return err
		}
	}

	return nil
}

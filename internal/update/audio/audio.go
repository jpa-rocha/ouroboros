// Package audio has the audio cmd internals
package audio

import (
	config "ouroboros/internal/config"
	update "ouroboros/internal/update"

	"github.com/spf13/cobra"
)

func RunAudio(cmd *cobra.Command, _ []string) error {
	ctx := cmd.Context()
	update.InstallPrereqs(ctx, false)
	update.InstallDriver(
		ctx,
		update.AUDIO,
		update.INSTALL_AUDIO,
		config.Opt.Repos.AudioGitRepo,
	)

	if config.Opt.Reboot.Yes {
		if err := update.RebootCmd(ctx); err != nil {
			return err
		}
	}

	return nil
}

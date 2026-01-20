// Package audio has the audio cmd internals
package audio

import (
	config "ouroboros/internal/config"
	update "ouroboros/internal/update"

	"github.com/spf13/cobra"
)

// RunAudio installs audio drivers by installing prerequisites and cloning/running the audio driver installation script.
// Takes a Cobra command and ignores the arguments.
// Optionally reboots the system if configured to do so.
// Returns error if installation or reboot fails.
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

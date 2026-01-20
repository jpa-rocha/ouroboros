// Package update selects which updates are needed
package update

import (
	"log/slog"

	all "ouroboros/cmd/update/all"
	audio "ouroboros/cmd/update/audio"
	bluetooth "ouroboros/cmd/update/bluetooth"
	camera "ouroboros/cmd/update/camera"

	"github.com/spf13/cobra"
)

// Cmd returns the update command with its subcommands for installing drivers.
// Provides options to update audio, bluetooth, camera drivers individually or all at once.
// Returns the configured update command.
func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Select which drivers to update.",
		Long:  "Allows the user to manually select which drivers to update.",
		Args:  cobra.NoArgs,
	}

	cmd.AddCommand(audio.Cmd())
	cmd.AddCommand(bluetooth.Cmd())
	cmd.AddCommand(camera.Cmd())
	cmd.AddCommand(all.Cmd())

	slog.Debug("command:", "value", cmd.Use)

	return cmd
}

// Package audio manages the audio command
package audio

import (
	"ouroboros/internal/update/audio"

	"github.com/spf13/cobra"
)

// Cmd returns the audio driver installation command.
// The command requires sudo privileges and reinstalls audio drivers after kernel updates.
// Returns the configured audio command.
func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "audio",
		Short: "Install audio drivers",
		Long:  "After kernel uppdates audio drivers need to be reinstalled [requires sudo]",
		Args:  cobra.NoArgs,
		RunE:  audio.RunAudio,
	}
	return cmd
}

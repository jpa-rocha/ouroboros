// Package audio manages the audio command
package audio

import (
	"ouroboros/internal/update/audio"

	"github.com/spf13/cobra"
)

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

// Package all manages the all command
package all

import (
	"ouroboros/internal/update/all"

	"github.com/spf13/cobra"
)

// Cmd returns the command to install all drivers.
// The command requires sudo privileges and reinstalls all drivers (audio, bluetooth, camera) after kernel updates.
// Returns the configured all command.
func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all",
		Short: "Install all drivers",
		Long:  "After kernel uppdates all drivers need to be reinstalled [requires sudo]",
		Args:  cobra.NoArgs,
		RunE:  all.RunAll,
	}
	return cmd
}

// Package camera manages the camera command
package camera

import (
	camera "ouroboros/internal/update/camera"

	"github.com/spf13/cobra"
)

// Cmd returns the camera driver installation command.
// The command requires sudo privileges and reinstalls camera firmware and drivers after kernel updates.
// Returns the configured camera command.
func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "camera",
		Short: "Install camera drivers",
		Long:  "After kernel uppdates camera drivers need to be reinstalled [requires sudo]",
		Args:  cobra.NoArgs,
		RunE:  camera.RunCamera,
	}
	return cmd
}

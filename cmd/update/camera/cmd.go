// Package camera manages the camera command
package camera

import (
	camera "ouroboros/internal/update/camera"

	"github.com/spf13/cobra"
)

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

// Package bluetooth manages the bluetooth command
package bluetooth

import (
	bluetooth "ouroboros/internal/update/bluetooth"

	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bluetooth",
		Short: "Install bluetooth drivers",
		Long:  "After kernel uppdates bluetooth drivers need to be reinstalled [requires sudo]",
		Args:  cobra.NoArgs,
		RunE:  bluetooth.RunBluetooth,
	}
	return cmd
}

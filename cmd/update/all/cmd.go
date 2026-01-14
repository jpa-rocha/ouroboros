// Package all manages the all command
package all

import (
	"ouroboros/internal/update/all"

	"github.com/spf13/cobra"
)

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

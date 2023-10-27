package cmd

import (
	"github.com/spf13/cobra"
)


func init() {
	var completion =  &cobra.Command{
		Use:       "completion [SHELL]",
		Short:     "Prints shell completion scripts",
		Long:      "This command prints shell completion code which must be evaluated to provide interactive completion",
		ValidArgs: []string{"bash", "zsh", "fish", "powershell"},
		Annotations: map[string]string{
			"commandType": "main",
		},
		Args: cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		RunE: func(cmd *cobra.Command, args []string) error {
			switch args[0] {
			case "bash":
				_ = cmd.Root().GenBashCompletion(cmd.OutOrStdout())
			case "zsh":
				_ = cmd.Root().GenZshCompletion(cmd.OutOrStdout())
			case "fish":
				_ = cmd.Root().GenFishCompletion(cmd.OutOrStdout(), true)
			case "powershell":
				_ = cmd.Root().GenPowerShellCompletion(cmd.OutOrStdout())
			}

			return nil
		},
	}
	rootCmd.AddCommand(completion)
}
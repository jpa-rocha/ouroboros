package cmd

import (
	"os"
	"fmt"
	"github.com/spf13/cobra"
)


var rootCmd = &cobra.Command{
	Use: "ouroboros",
	Short: "Ouroboros updates the audio, the bluetooth drivers",
	Long: "If running Ubunto on a MacBook, when the kernel is updated, ouroboros updates the audio, bluetooth or both drivers.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("----------- Ouroboros ------------")
		fmt.Println("Which drivers should be installed?")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
	  fmt.Fprintln(os.Stderr, err)
	  os.Exit(1)
	}
  }
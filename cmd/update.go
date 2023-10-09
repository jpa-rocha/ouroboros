package cmd

import (
	"os"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	var reboot bool
	var update = &cobra.Command {
		Use: "update",
		Short: "Select which drivers to update.",
		Long: "Allows the user to manually select which drivers to update.",
		Args: cobra.NoArgs,
	}

	var audio  = &cobra.Command {
		Use: "audio",
		Short: "Install audio drivers",
		Long: "After kernel uppdates audio drivers need to be reinstalled [requires sudo]",
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			installAudio()
			if reboot {
				
			}
		},
	}
	audio.Flags().BoolVarP(&reboot, "reboot", "r", viper.GetBool("restart"), "reboot system after updates")
	rootCmd.AddCommand(update)
}

func installAudio() {
	audioFolder := "audio/"
	install_cmd_audio := "./install.cirrus.driver.sh"
	ExecuteCommand(
		[]string{"git", "clone", viper.GetString("audio_git_repo"), audioFolder},
		"downloading needed repository...",
		"error: there was a problem downloading the needed files",
	)
	os.Chdir(audioFolder)
	ExecuteCommand(
		[]string{install_cmd_audio},
		"installing the audio drivers...",
		"error: there was a problem installing the drivers",
	)
	os.Chdir("../")
	ExecuteCommand(
		[]string{"rm", "-rf", audioFolder},
		"removing repository...",
		"error: there was a problem removing the repository",
	)
	fmt.Println("installed audio drivers successfuly")
}
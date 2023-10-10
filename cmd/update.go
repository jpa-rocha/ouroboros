package cmd

import (
	"fmt"
	"os"
	"os/exec"

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

	var audio = &cobra.Command {
		Use: "audio",
		Short: "Install audio drivers",
		Long: "After kernel uppdates audio drivers need to be reinstalled [requires sudo]",
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			installDriver(
				"audio/",
				"./install.cirrus.driver.sh",
				viper.GetString("audioGitRepo"),
			)
			if reboot {
				rebootCmd()
			}
		},
	}

	var bluetooth = &cobra.Command {
		Use: "bluetooth",
		Short: "Install bluetooth drivers",
		Long: "After kernel uppdates bluetooth drivers need to be reinstalled [requires sudo]",
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			installDriver(
				"bluetooth/",
				"./install.bluetooth.sh",
				viper.GetString("bluetoothGitRepo"),
			)
			if reboot {
				rebootCmd()
			}
		},
	}

	var everything = &cobra.Command {
		Use: "everything",
		Short: "Install all drivers",
		Long: "After kernel uppdates all drivers need to be reinstalled [requires sudo]",
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			installDriver(
				"audio/",
				"./install.cirrus.driver.sh",
				viper.GetString("audioGitRepo"),
			)
			installDriver(
				"bluetooth/",
				"./install.bluetooth.sh",
				viper.GetString("bluetoothGitRepo"),
			)
			if reboot {
				rebootCmd()
			}
		},
	} 

	audio.Flags().BoolVarP(&reboot, "reboot", "r", viper.GetBool("restart"), "reboot system after updates")
	bluetooth.Flags().BoolVarP(&reboot, "reboot", "r", viper.GetBool("restart"), "reboot system after updates")
	everything.Flags().BoolVarP(&reboot, "reboot", "r", viper.GetBool("restart"), "reboot system after updates")
	rootCmd.AddCommand(update)
	update.AddCommand(audio, bluetooth, everything)
}

func handleError(err error) {
	if err != nil {
		os.Exit(1)
	}
}

func rebootCmd() {
	reboot := exec.Command("sudo", "reboot")
	err := reboot.Run()
	if err != nil {
		fmt.Println("error: could not reboot")
	}
}

func installDriver(folder string, cmd string, repo string) {
	err := ExecuteCommand(
		[]string{"git", "clone", repo, folder},
		"downloading needed repository...",
		"error: there was a problem downloading the needed files",
	)
	handleError(err)
	os.Chdir(folder)
	err = ExecuteCommand(
		[]string{cmd},
		"installing the audio drivers...",
		"error: there was a problem installing the drivers",
	)
	handleError(err)
	os.Chdir("../")
	err = ExecuteCommand(
		[]string{"rm", "-rf", folder},
		"removing repository...",
		"error: there was a problem removing the repository",
	)
	handleError(err)
	fmt.Println("installed audio drivers successfuly")
}
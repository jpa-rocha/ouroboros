package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	update := &cobra.Command{
		Use:   "update",
		Short: "Select which drivers to update.",
		Long:  "Allows the user to manually select which drivers to update.",
		Args:  cobra.NoArgs,
	}

	var reboot bool
	audio := &cobra.Command{
		Use:   "audio",
		Short: "Install audio drivers",
		Long:  "After kernel uppdates audio drivers need to be reinstalled [requires sudo]",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			installPrereqs(false)
			installDriver(
				"audio",
				"./install.cirrus.driver.sh",
				viper.GetString("audioGitRepo"),
			)
			if reboot {
				rebootCmd()
			}
		},
	}

	bluetooth := &cobra.Command{
		Use:   "bluetooth",
		Short: "Install bluetooth drivers",
		Long:  "After kernel uppdates bluetooth drivers need to be reinstalled [requires sudo]",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			installPrereqs(false)
			installDriver(
				"bluetooth",
				"./install.bluetooth.sh",
				viper.GetString("bluetoothGitRepo"),
			)
			if reboot {
				rebootCmd()
			}
		},
	}

	camera := &cobra.Command{
		Use:   "camera",
		Short: "Install camera drivers",
		Long:  "After kernel uppdates camera drivers need to be reinstalled [requires sudo]",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			installPrereqs(true)
			installCamera()
			if reboot {
				rebootCmd()
			}
		},
	}

	everything := &cobra.Command{
		Use:   "everything",
		Short: "Install all drivers",
		Long:  "After kernel uppdates all drivers need to be reinstalled [requires sudo]",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			installPrereqs(true)
			installDriver(
				"audio",
				"./install.cirrus.driver.sh",
				viper.GetString("audioGitRepo"),
			)
			installDriver(
				"bluetooth",
				"./install.bluetooth.sh",
				viper.GetString("bluetoothGitRepo"),
			)
			installCamera()
			if reboot {
				rebootCmd()
			}
		},
	}

	update.PersistentFlags().BoolVarP(&reboot, "reboot", "r", viper.GetBool("reboot"), "reboot system after updates")
	rootCmd.AddCommand(update)
	update.AddCommand(audio, bluetooth, camera, everything)
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

func installDriver(driver string, cmd string, repo string) {
	fmt.Printf("starting %s driver installation...\n", driver)
	err := ExecuteCommand(
		[]string{"git", "clone", repo, driver},
		"downloading needed repository...",
		"error: there was a problem downloading the needed files",
	)
	handleError(err)
	err = os.Chdir(driver)
	handleError(err)
	err = ExecuteCommand(
		[]string{cmd},
		"installing the drivers...",
		"error: there was a problem installing the drivers",
	)
	handleError(err)
	os.Chdir("../")
	err = ExecuteCommand(
		[]string{"rm", "-rf", driver},
		"removing repository...",
		"error: there was a problem removing the repository",
	)
	handleError(err)
	fmt.Println("drivers installed successfuly")
}

func installPrereqs(isCamera bool) {
	ExecuteCommand(
		[]string{"pacman", "-Su"},
		"checking for updates...",
		"error: there was a problem checking for updates",
	)
	err := ExecuteCommand(
		[]string{"pacman", "-Sy", "gcc", "linux-headers-generic", "patch", "make", "wget"},
		"installing prerequisites...",
		"error: there was a problem installing prerequisites",
	)
	if isCamera {
		err = ExecuteCommand(
			[]string{"pacman", "-Sy", "facetimehd-dkms"},
			"installing prerequisites...",
			"error: there was a problem installing prerequisites",
		)
	}
	handleError(err)
}

func installCamera() {
	fmt.Println("drivers installed successfuly")
}

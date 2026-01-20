// Package camera has the internals for the camera command
package camera

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	config "ouroboros/internal/config"
	update "ouroboros/internal/update"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RunCamera installs camera firmware and drivers by installing prerequisites and running the camera installation process.
// Takes a Cobra command and ignores the arguments.
// Optionally reboots the system if configured to do so.
// Returns error if installation or reboot fails.
func RunCamera(cmd *cobra.Command, _ []string) error {
	ctx := cmd.Context()
	update.InstallPrereqs(ctx, true)
	InstallCamera(ctx)

	if config.Opt.Reboot.Yes {
		if err := update.RebootCmd(ctx); err != nil {
			return err
		}
	}

	return nil
}

// InstallCamera orchestrates installation of camera firmware and drivers.
// Takes context for logging. Clones and compiles camera firmware, then installs camera drivers.
// Cleans up cloned repositories after successful installation.
// Exits on any error encountered during the process.
func InstallCamera(ctx context.Context) {
	cameraFirmware := "cameraFirmware"
	cameraDrivers := "cameraDrivers"

	installCameraFirmware(ctx, cameraFirmware)
	installCameraDriver(ctx, cameraDrivers)
	err := update.ExecuteCommand(ctx,
		[]string{"rm", "-rf", cameraFirmware},
		"removing repository...",
		"error: there was a problem removing the repository",
	)
	update.HandleError(err)

	err = update.ExecuteCommand(ctx,
		[]string{"rm", "-rf", cameraDrivers},
		"removing repository...",
		"error: there was a problem removing the repository",
	)
	update.HandleError(err)
	fmt.Println("drivers installed successfuly")
}

// installCameraFirmware clones and compiles camera firmware from a configured repository.
// Takes context for logging and the directory name to clone into.
// Runs make and make install to compile and install firmware.
// Returns to parent directory after completion.
// Exits on any error encountered during the process.
func installCameraFirmware(ctx context.Context, cameraFirmware string) {
	fmt.Printf("starting %s firmware installation...\n", cameraFirmware)
	err := update.ExecuteCommand(ctx,
		[]string{"git", "clone", viper.GetString("cameraFirmware"), cameraFirmware},
		"downloading needed repository...",
		"error: there was a problem downloading the needed files",
	)

	update.HandleError(err)
	err = os.Chdir(cameraFirmware)
	update.HandleError(err)
	err = update.ExecuteCommand(
		ctx,
		[]string{"make"},
		"installing the drivers...",
		"error: there was a problem installing the drivers",
	)

	update.HandleError(err)
	err = update.ExecuteCommand(
		ctx,
		[]string{"make", "install"},
		"installing the drivers...",
		"error: there was a problem installing the drivers",
	)
	update.HandleError(err)

	if err := os.Chdir("../"); err != nil {
		slog.ErrorContext(ctx, "changing directories:", "error:", err.Error())
		update.HandleError(err)
	}
}

// installCameraDriver clones, compiles, and loads the camera driver kernel module.
// Takes context for logging and the directory name to clone into.
// Runs make, make install, depmod, and modprobe facetimehd to compile and load the driver.
// Returns to parent directory after completion.
// Exits on any error encountered during the process.
func installCameraDriver(ctx context.Context, cameraDrivers string) {
	err := update.ExecuteCommand(
		ctx,
		[]string{"git", "clone", viper.GetString("cameraDrivers"), cameraDrivers},
		"downloading needed repository...",
		"error: there was a problem downloading the needed files",
	)

	update.HandleError(err)
	err = os.Chdir(cameraDrivers)
	update.HandleError(err)

	err = update.ExecuteCommand(
		ctx,
		[]string{"make"},
		"installing the drivers...",
		"error: there was a problem installing the drivers",
	)

	update.HandleError(err)
	err = update.ExecuteCommand(
		ctx,
		[]string{"make", "install"},
		"installing the drivers...",
		"error: there was a problem installing the drivers",
	)

	update.HandleError(err)
	err = update.ExecuteCommand(
		ctx,
		[]string{"depmod"},
		"installing the drivers...",
		"error: there was a problem installing the drivers",
	)

	update.HandleError(err)
	err = update.ExecuteCommand(
		ctx,
		[]string{"modprobe", "facetimehd"},
		"installing the drivers...",
		"error: there was a problem installing the drivers",
	)

	update.HandleError(err)

	if err := os.Chdir("../"); err != nil {
		slog.ErrorContext(ctx, "changing directories:", "error:", err.Error())
		update.HandleError(err)
	}
}

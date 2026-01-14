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

// Package update has the internal utils for audio, bluetooth & camera commands
package update

import (
	"context"
	"fmt"
	"log/slog"
	"net/url"
	"os"
	"os/exec"
	"strings"
)

const (
	AUDIO     = "audio"
	BLUETOOTH = "bluetooth"
	CAMERA    = "camera"
	ALL       = "all"

	INSTALL_AUDIO     = "./install.cirrus.driver.sh"
	INSTALL_BLUETOOTH = "./install.bluetooth.sh"
)

func ExecuteCommand(
	ctx context.Context,
	command []string,
	successMsg string,
	errorMsg string,
) error {
	execCommand := exec.Command("sudo", command...)
	output, err := execCommand.Output()
	if err != nil {
		if len(string(output)) > 0 {
			slog.ErrorContext(ctx, "executing command:", "error:", err.Error())
			return err
		}
	} else {
		stringOutput := strings.TrimSpace(string(output))
		printString := strings.ToLower(string(stringOutput))
		slog.InfoContext(ctx, successMsg, "value:", printString)
	}
	return nil
}

func HandleError(err error) {
	if err != nil {
		os.Exit(1)
	}
}

func InstallDriver(ctx context.Context, driver string, cmd string, repo url.URL) {
	fmt.Printf("starting %s driver installation...\n", driver)
	err := ExecuteCommand(ctx,
		[]string{"git", "clone", repo.String(), driver},
		"downloading needed repository...",
		"error: there was a problem downloading the needed files",
	)
	HandleError(err)
	err = os.Chdir(driver)
	HandleError(err)
	err = ExecuteCommand(ctx,
		[]string{cmd},
		"installing the drivers...",
		"error: there was a problem installing the drivers",
	)
	HandleError(err)
	if err := os.Chdir("../"); err != nil {
		slog.ErrorContext(ctx, "changing directories:", "error:", err.Error())
		HandleError(err)
	}
	err = ExecuteCommand(ctx,
		[]string{"rm", "-rf", driver},
		"removing repository...",
		"error: there was a problem removing the repository",
	)
	HandleError(err)
	fmt.Println("drivers installed successfuly")
}

func InstallPrereqs(ctx context.Context, isCamera bool) {
	if err := ExecuteCommand(ctx,
		[]string{"apt", "update"},
		"checking for updates...",
		"error: there was a problem checking for updates",
	); err != nil {
		HandleError(err)
	}
	if err := ExecuteCommand(ctx,
		[]string{"apt", "upgrade", "-y"},
		"applying updates...",
		"error: there was a problem applying updates",
	); err != nil {
		HandleError(err)
	}
	if err := ExecuteCommand(ctx,
		[]string{"apt", "install", "-y", "wget", "make", "gcc", "linux-headers-generic"},
		"installing prerequisites...",
		"error: there was a problem installing prerequisites",
	); err != nil {
		HandleError(err)
	}
	if isCamera {
		if err := ExecuteCommand(
			ctx,
			[]string{
				"apt",
				"install",
				"-y",
				"wget",
				"xz-utils",
				"curl",
				"cpio",
				"git",
				"kmod",
				"libssl-dev",
				"checkinstall",
			},
			"installing prerequisites...",
			"error: there was a problem installing prerequisites",
		); err != nil {
			HandleError(err)
		}
	}
}

func RebootCmd(ctx context.Context) error {
	reboot := exec.Command("sudo", "reboot")
	err := reboot.Run()
	if err != nil {
		slog.ErrorContext(ctx, "running reboot command:", "error:", err.Error())
		return err
	}

	return nil
}

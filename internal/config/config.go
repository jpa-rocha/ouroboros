// Package config defines and populates the config struct
package config

import (
	"context"
	"log/slog"
	"net/url"
	"os"
	"strings"

	"github.com/spf13/viper"
)

const (
	HTTPTimeout = 5
)

type Config struct {
	Logger   Logger
	Settings Settings
	Repos    Repos
	Reboot   Reboot
	Secrets  Secrets
}

type Logger struct {
	Level string
}

type Settings struct {
	MTU     int
	Address string
}

type Repos struct {
	AudioGitRepo          url.URL
	BlueToothGitRepo      url.URL
	CameraFirmwareGitRepo url.URL
	CameraDriversGitRepo  url.URL
}

type Reboot struct {
	Yes bool
}

type Secrets struct {
	Path string
	VPN  VPN `yaml:"vpn"`
}

type VPN struct {
	IPSec IPSec `yaml:"ipsec"`
	ID    ID    `yaml:"id"`
}

type IPSec struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type ID struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

//nolint:gochecknoglobals
var Opt Config

const value = "value"

func InitConfig(parentContext context.Context) {
	viper.SetConfigName("ouroboros")
	viper.SetConfigType("toml")
	viper.AddConfigPath("$HOME/.config/ouroboros/")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("OUROBOROS")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		slog.ErrorContext(parentContext, "fatal error config file", slog.String(value, err.Error()))
		os.Exit(1)
	}

	audioRepo := viper.GetString("repos.audio_git_repo")
	audioURL, err := url.Parse(audioRepo)
	if err != nil {
		slog.ErrorContext(parentContext, "could not parse audio repo url", "error:", err.Error())
		os.Exit(1)
	}

	bluetoothRepo := viper.GetString("repos.bluetooth_git_repo")
	bluetoothURL, err := url.Parse(bluetoothRepo)
	if err != nil {
		slog.ErrorContext(
			parentContext,
			"could not parse bluetooth repo url",
			"error:",
			err.Error(),
		)
		os.Exit(1)
	}

	cameraFirmware := viper.GetString("repos.camera_firmware")
	cameraFirmwareURL, err := url.Parse(cameraFirmware)
	if err != nil {
		slog.ErrorContext(
			parentContext,
			"could not parse camere firmware repo url",
			"error:",
			err.Error(),
		)
		os.Exit(1)
	}

	cameraDrivers := viper.GetString("repos.camera_drivers")
	cameraDriversURL, err := url.Parse(cameraDrivers)
	if err != nil {
		slog.ErrorContext(
			parentContext,
			"could not parse camere drivers repo url",
			"error:",
			err.Error(),
		)
		os.Exit(1)
	}

	Opt = Config{
		Logger{
			Level: viper.GetString("logger.level"),
		},
		Settings{
			MTU:     viper.GetInt("settings.mtu"),
			Address: viper.GetString("settings.address"),
		},
		Repos{
			AudioGitRepo:          *audioURL,
			BlueToothGitRepo:      *bluetoothURL,
			CameraFirmwareGitRepo: *cameraFirmwareURL,
			CameraDriversGitRepo:  *cameraDriversURL,
		},
		Reboot{
			Yes: viper.GetBool("reboot.yes"),
		},
		Secrets{
			Path: viper.GetString("secrets.path"),
		},
	}

	defaults()
}

func defaults() {
	viper.SetDefault("logger.level", "INFO")
}

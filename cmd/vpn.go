package cmd

import (
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)


func init() {
	var port int
	var vpn = &cobra.Command {
		Use: "vpn",
		Short: "Manages connection to the preset vpn network",
		ValidArgs: []string{"connect", "disconnect"},
		Long: "After setting a vpn network using the vpnc tool ouroboros vpn calls the command to connect or disconnect from it [requires sudo]",
		Args: cobra.NoArgs,
	}
	
	var connect  = &cobra.Command {
		Use: "connect",
		Short: "Connects to the preset vpn network",
		Long: "After setting a vpn network using the vpnc tool, vpn connect calls the command to connect or disconnect from it [requires sudo]",
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			// Close any previous connections
			ExecuteCommand(
				[]string{"vpnc-disconnect"},
				"checking for existing connections...",
				"no vpnc found running",
			)
			time.Sleep(5 * time.Second)
			ExecuteCommand(
				[]string{"vpnc", viper.GetString("vpnName"), "--local-port", strconv.Itoa(port)},
				"success: you are connected to the vpn",
				"error: could not connect to vpn",
			)
		},
	}

	var disconnect  = &cobra.Command {
		Use: "disconnect",
		Short: "Disconnects from the preset vpn network",
		Long: "After setting a vpn network using the vpnc tool, vpn disconnect calls the command to connect or disconnect from it [requires sudo]",
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			ExecuteCommand(
				[]string{"vpnc-disconnect"},
				"success:",
				"no vpnc found running",
			)
		},
	}
	connect.Flags().IntVarP(&port, "port", "p", viper.GetInt("vpnPort"), "connect to vpn on selected port")
	vpn.AddCommand(connect, disconnect)
	rootCmd.AddCommand(vpn)
}

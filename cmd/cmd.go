// Package cmd sets up the main cmd and gathers all subcommands
package cmd

import (
	"context"
	"fmt"

	update "ouroboros/cmd/update"
	vpn "ouroboros/cmd/vpn"
	config "ouroboros/internal/config"
	logger "ouroboros/internal/logger"
	secrets "ouroboros/internal/secrets"

	"github.com/spf13/cobra"
)

func RootCmd() *cobra.Command {
	ctx := context.Background()
	cmd := cobra.Command{
		Use:   "ouroboros",
		Short: "Ouroboros automates needed tasks for running Ubuntu on a MacBook",
		Long:  "To keep the Ubuntu powered MacBooks functioning at work, Ouroboros is used to run commands that are often needed",
		PersistentPreRun: func(_ *cobra.Command, _ []string) {
			logger.Init()
			config.InitConfig(ctx)
			secrets.SecretsInit(ctx)
		},
	}
	cmd.SetContext(ctx)
	cmd.AddCommand(vpn.Cmd())
	cmd.AddCommand(update.Cmd())
	ctx, cancel := context.WithTimeout(cmd.Context(), config.HTTPTimeout)

	defer cancel()

	return &cmd
}

func Execute() error {
	ctx := RootCmd().Context()
	if err := RootCmd().ExecuteContext(ctx); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

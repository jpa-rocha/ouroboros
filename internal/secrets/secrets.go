// Package secrets gets secrets from sops file and adds them to global config
package secrets

import (
	"context"
	"log/slog"
	"os"

	config "ouroboros/internal/config"

	"github.com/getsops/sops/v3/decrypt"
	"gopkg.in/yaml.v3"
)

func SecretsInit(ctx context.Context) {
	secretsByteArray, err := decrypt.File(config.Opt.Secrets.Path, "yaml")
	if err != nil {
		slog.ErrorContext(
			ctx,
			"fatal error secrets file",
			slog.String(config.Opt.Secrets.Path, err.Error()),
		)
		os.Exit(1)
	}

	sopsSecrets := config.Secrets{
		Path: config.Opt.Secrets.Path,
	}

	err = yaml.Unmarshal(secretsByteArray, &sopsSecrets)
	if err != nil {
		slog.ErrorContext(ctx, "fatal error unmarshalling secrets", "status", err.Error())
		os.Exit(1)
	}

	config.Opt.Secrets = sopsSecrets
}

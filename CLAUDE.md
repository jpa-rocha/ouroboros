# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

**Ouroboros** is a Go CLI tool that automates system tasks for running Ubuntu on MacBooks. It manages VPN connections and installs/updates drivers for audio, Bluetooth, and camera hardware specific to MacBooks.

The project uses:
- **Go 1.25** for the application
- **Nix** for reproducible development environments and builds
- **Cobra** for CLI command structure
- **SOPS** for secrets management
- **Viper** for configuration management

## Development Commands

All commands are available in the development shell via `nix flake show` or by entering `nix develop`.

### Build
```bash
build
```
Compiles the binary: `go build -o ouroboros main.go`

### Tests
```bash
tests
```
Runs all tests with verbose output and fail-fast: `go test -v -failfast ./...`

To run a specific test:
```bash
go test -v -run TestName ./path/to/package
```

### Linting
```bash
lint
```
Runs all Nix flake checks, including Go linting via golangci-lint, formatting checks (gofumpt, golines), and security scanning (trufflehog, ripsecrets).

### Dependencies
```bash
tidy
```
Updates and vendors Go dependencies: `go mod tidy && go mod vendor`

## Project Structure

### `/cmd`
CLI command definitions using Cobra. Contains:
- **`cmd.go`**: Root command setup, initializes config, logger, and secrets before any subcommand runs
- **`vpn/`**: VPN start/stop commands
- **`update/`**: Driver update commands (all, audio, bluetooth, camera)

### `/internal`
Core business logic organized by feature:
- **`config/`**: Configuration parsing from `ouroboros.toml` via Viper. Loads git repository URLs and VPN settings. Config is available globally via `config.Opt` after `config.InitConfig()` is called.
- **`logger/`**: Structured logging setup using slog with tint formatting
- **`secrets/`**: Encrypted secrets management using SOPS, reads from `secrets.yaml`
- **`vpn/`**: VPN connection handling (start/stop commands)
- **`update/`**: Driver installation logic for audio, bluetooth, camera
- **`utils/`**: Utility functions and error handling

### `/nix`
Nix configuration for reproducible builds and development:
- **`scripts.nix`**: Defines build, test, lint, and tidy commands
- **`devshells.nix`**: Development environment with Go, golangci-lint, SOPS, age, git-lfs
- **`pre-commit.nix`**: Git hooks configuration (gofumpt, golangci-lint, prettier, trufflehog, ripsecrets)
- **`ouroboros.nix`**: Build configuration for the CLI binary
- **`treefmt.nix`**: Code formatting configuration

## Architecture & Key Patterns

### Configuration Flow
1. **Initialization**: `RootCmd()` in `cmd.go` calls `config.InitConfig()`, `logger.Init()`, and `secrets.SecretsInit()` in `PersistentPreRun`
2. **Config Source**: Looks for `ouroboros.toml` in `$HOME/.config/ouroboros/` or current directory. Environment variables override config (prefix: `OUROBOROS_`)
3. **Global State**: Configuration is stored in global `config.Opt` after initialization and accessed throughout the app
4. **Secrets**: Encrypted YAML file path is specified in config, loaded on startup

### Command Structure
Each command (VPN, Update) follows the pattern:
- Top-level package defines command structure
- Sub-packages define individual subcommands
- Implementation logic in `/internal` called from command handlers
- Context passed through for logging and cancellation

### Secrets Management
- Uses SOPS with age encryption
- Secrets stored in `secrets.yaml` (encrypted)
- Path configured in `ouroboros.toml` under `[secrets]`
- VPN credentials (IPSec + ID) stored in YAML structure

## Git Hooks & Pre-commit

The project uses `git-hooks.nix` for automated checks:
- **Go**: gofumpt formatting, golangci-lint linting, golines line length
- **Nix**: nixfmt-rfc-style formatting, deadnix dead code detection
- **Secrets**: trufflehog, ripsecrets detection
- **Other**: YAML validation, shell script checks, prettier formatting
- **Custom**: `create-version` post-commit hook for version tagging

## Key Dependencies

- **github.com/spf13/cobra**: CLI framework
- **github.com/spf13/viper**: Configuration management
- **github.com/getsops/sops/v3**: Secrets encryption/decryption
- **github.com/lmittmann/tint**: Structured logging formatting
- **gopkg.in/yaml.v3**: YAML parsing for config/secrets

## Important Files

- **`main.go`**: Entry point, calls `cmd.Execute()`
- **`cmd/cmd.go`**: Root command initialization (line 17-36)
- **`internal/config/config.go`**: Config struct definitions and parsing (line 18-65 for types, line 71-151 for initialization)
- **`ouroboros.toml`**: Main configuration file with repos, VPN address, logging level
- **`.sops.yaml`**: SOPS configuration for encryption keys

## Development Notes

- Code uses context extensively for cancellation and logging
- Errors are wrapped with `fmt.Errorf("%w", err)` for stack trace preservation
- All logging goes through structured slog logger (initialized in RootCmd PersistentPreRun)
- Nix flake handles all tool versioning and environment consistency
- Pre-commit hooks run automatically; use `git commit --no-verify` to bypass if needed

{
  inputs = {
    flake-utils.url = "github:numtide/flake-utils";
    nix-filter.url = "github:numtide/nix-filter";
    nixpkgs-stable.url = "github:nixos/nixpkgs/nixos-25.05";
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    pre-commit-hooks.url = "github:cachix/git-hooks.nix";
    treefmt-nix.url = "github:numtide/treefmt-nix";
  };

  outputs =
    { self, ... }@inputs:
    inputs.flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = inputs.nixpkgs.legacyPackages.${system};
        lib = pkgs.lib;

        inherit (lib) makeBinPath;

        filterOvr = inputs.nix-filter.lib;

        go = pkgs.go_1_25;
        goBuild = pkgs.buildGo125Module;

        golangci-lint = pkgs.callPackage ./nix/golangci-lint.nix {
          inherit go goBuild makeBinPath;
        };

        scripts = import ./nix/scripts.nix {
          inherit pkgs go inputs;
        };

        treefmtEval = inputs.treefmt-nix.lib.evalModule pkgs ./nix/treefmt.nix;

        preCommit = import ./nix/pre-commit.nix {
          inherit golangci-lint;
          inherit (scripts) create-version;
          inherit (pkgs)
            gofumpt
            nixfmt-rfc-style
            prettier
            trufflehog
            ;
        };

        ouroboros = pkgs.callPackage ./nix/ouroboros.nix {
          inherit goBuild filterOvr;
        };
      in
      {
        devShells = import ./nix/devshells.nix {
          inherit
            go
            golangci-lint
            ouroboros
            scripts
            self
            system
            ;
          inherit (pkgs)
            age
            git
            git-lfs
            golangci-lint-langserver
            gopls
            mkShell
            sops
            ;
        };

        checks = import ./nix/checks.nix {
          inherit
            inputs
            preCommit
            system
            treefmtEval
            ;
        };

        formatter = treefmtEval.config.build.wrapper;
      }
    );
}

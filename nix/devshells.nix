{
  age,
  git,
  git-lfs,
  go,
  golangci-lint,
  golangci-lint-langserver,
  gopls,
  ouroboros,
  mkShell,
  scripts,
  self,
  system,
  ...
}:
{
  default = mkShell {
    packages = [
      age
      git
      git-lfs
      go
      golangci-lint
      golangci-lint-langserver
      gopls
      ouroboros
      scripts.create-version
      scripts.lint
      scripts.build
      scripts.test
      scripts.tidy
    ];
    inherit (self.checks.${system}.pre-commit-check) shellHook;
  };
}

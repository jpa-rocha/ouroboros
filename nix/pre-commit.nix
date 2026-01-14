{
  create-version,
  gofumpt,
  golangci-lint,
  nixfmt-rfc-style,
  prettier,
  trufflehog,
  ...
}:
{
  src = ./.;
  excludes = [
    "flake.lock"
    "vendor/.+"
    "secrets_gateway.yaml"
    "secrets_vm.yaml"
    ".envrc"
    "keys/.+"
  ];
  hooks = {
    check-yaml.enable = true;
    convco.enable = true;
    deadnix.enable = true;
    ripsecrets.enable = true;
    shellcheck.enable = true;

    nixfmt-rfc-style = {
      enable = true;
      package = nixfmt-rfc-style;
    };

    trufflehog = {
      enable = true;
      package = trufflehog;
    };

    prettier = {
      enable = true;
      package = prettier;
      excludes = [
        "assets/ui/.+"
      ];
    };

    # typos = {
    #   enable = true;
    # };

    golines = {
      enable = true;
      settings.flags = "--base-formatter=${gofumpt}/bin/gofumpt";
    };

    golangci-lint = {
      enable = true;
      package = golangci-lint;
    };

    create-version = {
      enable = true;
      name = "create-version";
      extraPackages = [
        create-version
      ];
      entry = "create-version";
      language = "system";
      stages = [ "post-commit" ];
      always_run = true;
    };
  };
}

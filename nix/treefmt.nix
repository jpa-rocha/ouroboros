{ pkgs, ... }:
{
  projectRootFile = "flake.nix";

  programs = {
    deadnix.enable = true;
    golines.enable = true;
    nixfmt.enable = true;
    prettier.enable = true;
    templ.enable = true;

    shfmt = {
      enable = true;
      indent_size = 0;
    };
  };

  settings = {
    global.excludes = [
      "vendor/*"
      "Taskfile.yml"
      "secrets_gateway.yaml"
      "secrets_vm.yaml"
    ];

    formatter = {
      golines = {
        options = [
          "--base-formatter=${pkgs.gofumpt}/bin/gofumpt"
        ];
      };

      prettier = {
        package = pkgs.prettier;
        excludes = [
        ];
      };
    };
  };
}

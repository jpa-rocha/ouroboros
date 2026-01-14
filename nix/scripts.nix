{
  pkgs,
  go,
  ...
}:
{
  tidy = pkgs.writeShellScriptBin "tidy" ''
    set -euo pipefail
    export GOPRIVATE=git.wobcom.de
    ${go}/bin/go mod tidy
    ${go}/bin/go mod vendor
  '';

  test = pkgs.writeShellScriptBin "tests" ''
    set -euo pipefail
    ${go}/bin/go test -v -failfast ./...
  '';

  lint = pkgs.writeShellScriptBin "lint" ''
    set -euo pipefail
    ${pkgs.nix}/bin/nix flake check
  '';

  build = pkgs.writeShellScriptBin "build" ''
    set -euo pipefail
    ${go}/bin/go build -o ouroboros main.go
  '';

  create-version = pkgs.writeShellScriptBin "create-version" ''
    set -euo pipefail

    # Only show most recent tag without trailing commit information
    ${pkgs.git}/bin/git describe --tags | ${pkgs.gawk}/bin/awk "{split(\$0,a,\"-\"); print a[1];}" >VERSION.tmp

    # Only proceed if version number has actually changed (i.e. a new tag has been created)
    if ! ${pkgs.diffutils}/bin/cmp --silent VERSION.tmp VERSION; then
      NEWVER=$(cat VERSION)
      echo Adding tag "$NEWVER"
      ${pkgs.git}/bin/git tag -a "$NEWVER" -m ""
    fi

    rm VERSION.tmp
  '';
}

{ goBuild, filterOvr, ... }:
let
  srcFilter = filterOvr {
    root = ./.;
    include = [
      ./cmd
      ./go.mod
      ./go.sum
      ./internal
      ./main.go
      ./vendor
    ];
  };
in
goBuild {
  name = "ouroboros";
  src = srcFilter;
  doCheck = false;
  vendorHash = null;
  ldflags = [
    "-s"
    "-w"
    "-extldflags '-static'"
  ];
  env.CGO_ENABLED = 0;
}

{
  inputs,
  preCommit,
  system,
  treefmtEval,
  ...
}:
{
  formatting = treefmtEval.config.build.check inputs.self;
  pre-commit-check = inputs.pre-commit-hooks.lib.${system}.run preCommit;
}

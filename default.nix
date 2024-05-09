{ pkgs ? import
    (fetchTarball {
      name = "jpetrucciani-2024-05-08";
      url = "https://github.com/jpetrucciani/nix/archive/4d04f8b2747475903513bd31aa0fba2d9cf697db.tar.gz";
      sha256 = "1p696dpyq7whbfz4ps3mk5i9qycd9k1fjg1dzsv6r7dhbbkc5hvz";
    })
    { }
}:
let
  name = "ahab";

  tools = with pkgs; {
    cli = [
      coreutils
      nixpkgs-fmt
    ];
    go = [
      go
      go-tools
      gopls
    ];
    scripts = pkgs.lib.attrsets.attrValues scripts;
  };

  scripts = with pkgs; { };
  paths = pkgs.lib.flatten [ (builtins.attrValues tools) ];
  env = pkgs.buildEnv {
    inherit name paths; buildInputs = paths;
  };
in
(env.overrideAttrs (_: {
  inherit name;
  NIXUP = "0.0.6";
})) // { inherit scripts; }

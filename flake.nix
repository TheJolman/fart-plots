{
  description = "fart plots dev shell and build definition";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = {
    self,
    nixpkgs,
    flake-utils,
  }:
    flake-utils.lib.eachDefaultSystem (
      system: let
        pkgs = nixpkgs.legacyPackages.${system};
      in {
        # TODO: add pre-commit and default package here
        devShells.default = pkgs.mkShell {
          packages = with pkgs; [
            go
            go-tools
          ];

          shellHook = ''
            echo "Loaded dev shell."
          '';
        };
      }
    );
}

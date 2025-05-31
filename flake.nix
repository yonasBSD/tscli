{
  description = "Flake for github:jaxxstorm/tscli";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-25.05";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs =
    {
      self,
      nixpkgs,
      flake-utils,
    }:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = nixpkgs.legacyPackages.${system};

        tscli = pkgs.buildGoModule {
          pname = "tscli";
          version = "0.0.4";
          src = self;
          vendorHash = "sha256-KEI7rdYWz7XHU3tBTxGZ5HhemlSEy3Hh633qEG89Dv4=";
        };
      in
      {
        devShells.default = pkgs.mkShell {
          packages = with pkgs; [
            go
            gotools
            gopls
          ];
        };

        packages = {
          default = tscli;
          tscli = tscli;
          docker = pkgs.dockerTools.buildImage {
            name = "ghcr.io/jaxxstorm/tscli";
            tag = "latest";
            copyToRoot = tscli;
            created = self.lastModifiedDate;
          };
        };
      }
    );
}

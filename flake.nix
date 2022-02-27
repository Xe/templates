{
  description = "Xe's Nix Flake Templates";

  outputs = { self, nixpkgs }: {
    templates = {
      go-web-server = {
        path = ./go-web-server;
        description = "A basic Go web server setup";
      };
    };
  };
}

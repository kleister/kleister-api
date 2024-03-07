{
  description = "Nix flake for development";

  inputs = {
    nixpkgs = {
      url = "github:nixos/nixpkgs/nixpkgs-unstable";
    };

    devenv = {
      url = "github:cachix/devenv";
    };

    flake-parts = {
      url = "github:hercules-ci/flake-parts";
    };
  };

  outputs = inputs@{ flake-parts, ... }:
    flake-parts.lib.mkFlake { inherit inputs; } {
      imports = [
        inputs.devenv.flakeModule
      ];

      systems = [
        "x86_64-linux"
        "aarch64-linux"
        "x86_64-darwin"
        "aarch64-darwin"
      ];

      perSystem = { config, self', inputs', pkgs, system, ... }: {
        imports = [
          {
            _module.args.pkgs = import inputs.nixpkgs {
              inherit system;
              config.allowUnfree = true;
            };
          }
        ];

        devenv = {
          shells = {
            default = {
              name = "kleister-api";

              certificates = [
                "localhost"
                "kleister.local"
              ];

              hosts = {
                "kleister.local" = "127.0.0.1";
              };

              languages = {
                go = {
                  enable = true;
                  package = pkgs.go_1_22;
                };
              };

              services = {
                minio = {
                  enable = true;
                  accessKey = "9VKV3OI56N1077Y9IALV";
                  secretKey = "bwcRkW5w6uF6BWBqotsnMbwZSIDKQopy9DSo90ab";
                  buckets = [
                    "kleister"
                  ];
                };
                mysql = {
                  enable = true;
                  ensureUsers = [
                    {
                      name = "kleister";
                      password = "p455w0rd";
                      ensurePermissions = {
                        "kleister.*" = "ALL PRIVILEGES";
                      };
                    }
                  ];
                  initialDatabases = [{
                    name = "kleister";
                  }];
                };
                postgres = {
                  enable = true;
                  listen_addresses = "127.0.0.1";
                  initialScript = ''
                    CREATE USER kleister WITH ENCRYPTED PASSWORD 'p455w0rd';
                    GRANT ALL PRIVILEGES ON DATABASE kleister TO kleister;
                  '';
                  initialDatabases = [{
                    name = "kleister";
                  }];
                };
              };

              packages = with pkgs; [
                bingo
                gnumake
                httpie
                nixpkgs-fmt
                sqlite
              ];

              env = {
                CGO_ENABLED = "0";

                KLEISTER_API_LOG_LEVEL = "debug";
                KLEISTER_API_SERVER_PPROF = "true";

                # KLEISTER_API_SERVER_CERT = ".devenv/state/mkcert/localhost+1.pem";
                # KLEISTER_API_SERVER_KEY = ".devenv/state/mkcert/localhost+1-key.pem";

                KLEISTER_API_DATABASE_DRIVER = "sqlite3";
                KLEISTER_API_DATABASE_NAME = "storage/kleister.sqlite3";

                # KLEISTER_API_DATABASE_DRIVER = "mysql";
                # KLEISTER_API_DATABASE_ADDRESS = "127.0.0.1";
                # KLEISTER_API_DATABASE_PORT = "3306";
                # KLEISTER_API_DATABASE_USERNAME = "kleister";
                # KLEISTER_API_DATABASE_PASSWORD = "p455w0rd";
                # KLEISTER_API_DATABASE_NAME = "kleister";

                # KLEISTER_API_DATABASE_DRIVER = "postgres";
                # KLEISTER_API_DATABASE_ADDRESS = "127.0.0.1";
                # KLEISTER_API_DATABASE_PORT = "5432";
                # KLEISTER_API_DATABASE_USERNAME = "kleister";
                # KLEISTER_API_DATABASE_PASSWORD = "p455w0rd";
                # KLEISTER_API_DATABASE_NAME = "kleister";

                KLEISTER_API_UPLOAD_DRIVER = "file";
                KLEISTER_API_UPLOAD_PATH = "storage/uploads/";

                # KLEISTER_API_UPLOAD_DRIVER = "s3";
                # KLEISTER_API_UPLOAD_ENDPOINT = "127.0.0.1:9000";
                # KLEISTER_API_UPLOAD_BUCKET = "kleister";
                # KLEISTER_API_UPLOAD_REGION = "us-east-1";
                # KLEISTER_API_UPLOAD_ACCESS = "9VKV3OI56N1077Y9IALV";
                # KLEISTER_API_UPLOAD_SECRET = "bwcRkW5w6uF6BWBqotsnMbwZSIDKQopy9DSo90ab";

                KLEISTER_API_AUTH_CONFIG = "config/auth.yml";
                KLEISTER_API_ADMIN_USERS = "tboerger";
              };
            };
          };
        };
      };
    };
}

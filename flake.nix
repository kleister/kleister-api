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

    git-hooks = {
      url = "github:cachix/git-hooks.nix";
    };
  };

  outputs =
    inputs@{ flake-parts, ... }:
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

      perSystem =
        {
          config,
          self',
          inputs',
          pkgs,
          system,
          ...
        }:
        {
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

                git-hooks = {
                  hooks = {
                    nixfmt-rfc-style = {
                      enable = true;
                    };

                    gofmt = {
                      enable = true;
                    };

                    golangci-lint = {
                      enable = true;
                      entry = "go tool github.com/golangci/golangci-lint/cmd/golangci-lint run ./...";
                      pass_filenames = false;
                    };
                  };
                };

                languages = {
                  go = {
                    enable = true;
                    package = pkgs.go_1_24;
                  };
                };

                packages = with pkgs; [
                  cosign
                  gnumake
                  goreleaser
                  httpie
                  nixfmt-rfc-style
                  posting
                  sqlite
                  yq-go
                ];

                env = {
                  CGO_ENABLED = "0";

                  KLEISTER_API_LOG_LEVEL = "debug";
                  KLEISTER_API_LOG_PRETTY = "true";
                  KLEISTER_API_LOG_COLOR = "true";

                  KLEISTER_API_TOKEN_SECRET = "TxHrYxMAg01rBeEWrHn1BjOP";
                  KLEISTER_API_TOKEN_EXPIRE = "1h";

                  KLEISTER_API_SERVER_FRONTEND = "http://localhost:5173";
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
                  # KLEISTER_API_UPLOAD_ENDPOINT = "http://127.0.0.1:9000";
                  # KLEISTER_API_UPLOAD_PROXY = "true";
                  # KLEISTER_API_UPLOAD_PATHSTYLE = "true";
                  # KLEISTER_API_UPLOAD_PATH = "";
                  # KLEISTER_API_UPLOAD_BUCKET = "kleister";
                  # KLEISTER_API_UPLOAD_REGION = "us-east-1";
                  # KLEISTER_API_UPLOAD_ACCESS = "9VKV2OI56N1077Y9IALV";
                  # KLEISTER_API_UPLOAD_SECRET = "bwcRkW5w6uF6CWBqotsnMbwZSIDKQopy9DSo90ab";

                  KLEISTER_API_CLEANUP_ENABLED = "true";
                  KLEISTER_API_CLEANUP_INTERVAL = "5m";

                  KLEISTER_API_ADMIN_USERNAME = "admin";
                  KLEISTER_API_ADMIN_PASSWORD = "p455w0rd";
                  KLEISTER_API_ADMIN_EMAIL = "kleister@webhippie.de";
                };

                services = {
                  minio = {
                    enable = true;
                    accessKey = "9VKV2OI56N1077Y9IALV";
                    secretKey = "bwcRkW5w6uF6CWBqotsnMbwZSIDKQopy9DSo90ab";
                    buckets = [
                      "kleister"
                    ];
                  };
                  postgres = {
                    enable = true;
                    listen_addresses = "127.0.0.1";
                    initialScript = ''
                      CREATE USER kleister WITH ENCRYPTED PASSWORD 'p455w0rd';
                      GRANT ALL PRIVILEGES ON DATABASE kleister TO kleister;
                    '';
                    initialDatabases = [
                      {
                        name = "kleister";
                      }
                    ];
                  };
                };

                processes = {
                  kleister-server = {
                    exec = "make watch";

                    process-compose = {
                      readiness_probe = {
                        exec.command = "${pkgs.curl}/bin/curl -sSf http://localhost:8000/readyz";
                        initial_delay_seconds = 2;
                        period_seconds = 10;
                        timeout_seconds = 4;
                        success_threshold = 1;
                        failure_threshold = 5;
                      };

                      availability = {
                        restart = "on_failure";
                      };
                    };
                  };

                  minio = {
                    process-compose = {
                      readiness_probe = {
                        exec.command = "${pkgs.curl}/bin/curl -sSf http://localhost:9000/minio/health/live";
                        initial_delay_seconds = 2;
                        period_seconds = 10;
                        timeout_seconds = 4;
                        success_threshold = 1;
                        failure_threshold = 5;
                      };

                      availability = {
                        restart = "on_failure";
                      };
                    };
                  };
                };
              };
            };
          };
        };
    };
}

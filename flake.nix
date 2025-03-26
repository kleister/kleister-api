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
        inputs.git-hooks.flakeModule
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
                  javascript = {
                    enable = true;
                    package = pkgs.nodejs_20;
                  };
                };

                packages = with pkgs; [
                  go-task
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

                  KLEISTER_API_TOKEN_SECRET = "hgJKKJrSzI8pOxCjCnJHvNvK";
                  KLEISTER_API_TOKEN_EXPIRE = "1h";

                  KLEISTER_API_DATABASE_DRIVER = "sqlite3";
                  KLEISTER_API_DATABASE_NAME = "storage/kleister.sqlite3";

                  KLEISTER_API_UPLOAD_DRIVER = "file";
                  KLEISTER_API_UPLOAD_PATH = "storage/uploads/";

                  KLEISTER_API_CLEANUP_ENABLED = "true";
                  KLEISTER_API_CLEANUP_INTERVAL = "5m";

                  KLEISTER_API_ADMIN_USERNAME = "admin";
                  KLEISTER_API_ADMIN_PASSWORD = "p455w0rd";
                  KLEISTER_API_ADMIN_EMAIL = "kleister@webhippie.de";
                };

                services = {
                  minio = {
                    enable = true;
                    accessKey = "2p9FzWqUyhLrwiHpdviD";
                    secretKey = "IpYoERy9493E1wIsfXk9wCv3vHS8o6nMf101BDpT";
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
                    exec = "task watch:server";

                    process-compose = {
                      environment = [
                        "KLEISTER_API_SERVER_HOST=http://localhost:5173"
                      ];

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

                  kleister-webui = {
                    exec = "task watch:frontend";

                    process-compose = {
                      readiness_probe = {
                        exec.command = "${pkgs.curl}/bin/curl -sSf http://localhost:5173";
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

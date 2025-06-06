# Kleister: API server

[![General Workflow](https://github.com/kleister/kleister-api/actions/workflows/general.yml/badge.svg)](https://github.com/kleister/kleister-api/actions/workflows/general.yml) [![Join the Matrix chat at https://matrix.to/#/#kleister:matrix.org](https://img.shields.io/badge/matrix-%23kleister-7bc9a4.svg)](https://matrix.to/#/#kleister:matrix.org) [![Codacy Badge](https://app.codacy.com/project/badge/Grade/c4d0c564f786486c93e37d62db312746)](https://app.codacy.com/gh/kleister/kleister-api/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade) [![Go Reference](https://pkg.go.dev/badge/github.com/kleister/kleister-api.svg)](https://pkg.go.dev/github.com/kleister/kleister-api) [![GitHub Repo](https://img.shields.io/badge/github-repo-yellowgreen)](https://github.com/kleister/kleister-api)

> [!CAUTION]
> This project is in active development and does not provide any stable release
> yet, you can expect breaking changes until our first real release!

Kleister is a web UI to manage mod packs for the Minecraft, initially focused on
the Technic Launcher and MCUpdater. Even if there is an upstream version
available the Technic Launcher at [TechnicPack/TechnicSolder][solder] I prefered
to implement it in Go for the API and VueJS for the UI including some further
features like uploading the mods I want to manage and even generating docker
images directly out of the managed packs. Hosting Minecraft servers based on
docker images works pretty cool.

## Install

You can download prebuilt binaries from the [GitHub releases][releases] or from
our [download site][downloads]. Besides that we also prepared repositories for
DEB and RPM packages which can be found at [Baltorepo][baltorepo]. If you prefer
to use containers you could use our images published on [GHCR][ghcr],
[Docker Hub][dockerhub] or [Quay][quay]. You are a Mac user? Just take a look
at our [homebrew formula][homebrew]. If you need further guidance how to
install this take a look at our [documentation][docs].

## Build

If you are not familiar with [Nix][nix] it is up to you to have a working
environment for Go (>= 1.24.0) and Nodejs (22.x) as the setup won't we covered
within this guide. Please follow the official install instructions for
[Go][golang] and [Nodejs][nodejs]. Beside that we are using [go-task][gotask] to
define all commands to build this project.

```console
git clone https://github.com/kleister/kleister-api.git
cd kleister-api

task fe:install fe:build be:build
./bin/kleister-api -h
```

If you got [Nix][nix] and [Direnv][direnv] configured you can simply execute
the following commands to get al dependencies including [go-task][gotask] and
the required runtimes installed. You are also able to directly use the process
manager of [devenv][devenv]:

```console
cat << EOF > .envrc
use flake . --impure --extra-experimental-features nix-command
EOF

direnv allow
```

We are embedding all the static assets into the binary so there is no need for
any webserver or anything else beside launching this binary.

## Development

To start developing on this project you have to execute only a few commands. To
start development just execute those commands in different terminals:

```console
task watch:server
task watch:frontend
```

The development server of the backend should be running on
[http://localhost:8080](http://localhost:8080) while the frontend should be
running on [http://localhost:5173](http://localhost:5173). Generally it supports
hot reloading which means the services are automatically restarted/reloaded on
code changes.

If you got [Nix][nix] configured you can simply execute the [devenv][devenv]
command to start the frontend, backend, MariaDB, PostgreSQL and Minio:

```console
devenv up
```

## Security

If you find a security issue please contact
[kleister@webhippie.de](mailto:kleister@webhippie.de) first.

## Contributing

Fork -> Patch -> Push -> Pull Request

## Authors

-   [Thomas Boerger](https://github.com/tboerger)

## License

Apache-2.0

## Copyright

```console
Copyright (c) 2018 Thomas Boerger <thomas@webhippie.de>
```

[releases]: https://github.com/kleister/kleister-api/releases
[downloads]: https://dl.kleister.eu
[baltorepo]: https://kleister.baltorepo.com/stable/
[homebrew]: https://github.com/kleister/homebrew-kleister
[ghcr]: https://github.com/orgs/kleister/packages
[dockerhub]: https://hub.docker.com/r/kleister/kleister-api/tags/
[quay]: https://quay.io/repository/kleister/kleister-api?tab=tags
[docs]: https://kleister.eu/
[nix]: https://nixos.org/
[golang]: http://golang.org/doc/install.html
[gotask]: https://taskfile.dev/installation/
[direnv]: https://direnv.net/
[devenv]: https://devenv.sh/
[solder]: https://github.com/TechnicPack/TechnicSolder

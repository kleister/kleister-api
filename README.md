# Kleister: API server

[![General Workflow](https://github.com/kleister/kleister-api/actions/workflows/general.yml/badge.svg)](https://github.com/kleister/kleister-api/actions/workflows/general.yml) [![Join the Matrix chat at https://matrix.to/#/#kleister:matrix.org](https://img.shields.io/badge/matrix-%23kleister-7bc9a4.svg)](https://matrix.to/#/#kleister:matrix.org) [![Codacy Badge](https://app.codacy.com/project/badge/Grade/c4d0c564f786486c93e37d62db312746)](https://app.codacy.com/gh/kleister/kleister-api/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade) [![Go Reference](https://pkg.go.dev/badge/github.com/kleister/kleister-api.svg)](https://pkg.go.dev/github.com/kleister/kleister-api) [![GitHub Repo](https://img.shields.io/badge/github-repo-yellowgreen)](https://github.com/kleister/kleister-api)

Kleister is a web UI to manage mod packs for the Minecraft, initially focused on
the Technic Launcher and MCUpdater. Even if there is an upstream version
available the Technic Launcher at [TechnicPack/TechnicSolder][solder] I prefered
to implement it in Go for the API and VueJS for the UI including some further
features like uploading the mods I want to manage and even generating docker
images directly out of the managed packs. Hosting Minecraft servers based on
docker images works pretty cool.

## Install

You can download prebuilt binaries from the [GitHub releases][releases] or from
our [download site][downloads]. If you prefer to use containers you could use
our images published on [Docker Hub][dockerhub] or [Quay][quay]. You are a Mac
user? Just take a look at our [homebrew formula][homebrew]. If you need further
guidance how to install this take a look at our [documentation][docs].

## Development

Make sure you have a working Go environment, for further reference or a guide
take a look at the [install instructions][golang]. This project requires
Go >= v1.23.1, at least that's the version we are using.

```console
git clone https://github.com/kleister/kleister-api.git
cd kleister-api

make generate build

./bin/kleister-api -h
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
[downloads]: https://dl.kleister.eu/api
[homebrew]: https://github.com/kleister/homebrew-kleister
[dockerhub]: https://hub.docker.com/r/kleister/kleister-api/tags/
[quay]: https://quay.io/repository/kleister/kleister-api?tab=tags
[docs]: https://kleister.eu/
[golang]: http://golang.org/doc/install.html
[solder]: https://github.com/TechnicPack/TechnicSolder

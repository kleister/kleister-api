# Kleister: API server

[![Build Status](https://github.com/kleister/kleister-api/actions/workflows/general.yml/badge.svg)](https://github.com/kleister/kleister-api/actions) [![Join the Matrix chat at https://matrix.to/#/#kleister:matrix.org](https://img.shields.io/badge/matrix-%23kleister-7bc9a4.svg)](https://matrix.to/#/#kleister:matrix.org) [![Go Reference](https://pkg.go.dev/badge/github.com/kleister/kleister-api.svg)](https://pkg.go.dev/github.com/kleister/kleister-api) [![Go Report Card](https://goreportcard.com/badge/github.com/kleister/kleister-api)](https://goreportcard.com/report/github.com/kleister/kleister-api) [![Codacy Badge](https://app.codacy.com/project/badge/Grade/c4d0c564f786486c93e37d62db312746)](https://www.codacy.com/gh/kleister/kleister-api/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=kleister/kleister-api&amp;utm_campaign=Badge_Grade)

Kleister is a web UI to manage mod packs for the Minecraft, initially focused on
the Technic Launcher and MCUpdater. Even if there is an upstream version
available the Technic Launcher at [TechnicPack/TechnicSolder][solder] I prefered
to implement it in Go for the API and VueJS for the UI including some further
features like uploading the mods I want to manage and even generating docker
images directly out of the managed packs. Hosting Minecraft servers based on
docker images works pretty cool.

## Install

You can download prebuilt binaries from our [GitHub releases][releases], or you
can use our Docker images published on [Docker Hub][dockerhub] or [Quay][quay].
If you need further guidance how to install this take a look at our
[documentation][docs].

## Development

Make sure you have a working Go environment, for further reference or a guide
take a look at the [install instructions][golang]. This project requires
Go >= v1.17, at least that's the version we are using.

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
[dockerhub]: https://hub.docker.com/r/kleister/kleister-api/tags/
[quay]: https://quay.io/repository/kleister/kleister-api?tab=tags
[docs]: https://kleister.eu/
[golang]: http://golang.org/doc/install.html
[solder]: https://github.com/TechnicPack/TechnicSolder

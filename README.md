# Kleister: API server

[![Build Status](http://github.dronehippie.de/api/badges/kleister/kleister-api/status.svg)](http://github.dronehippie.de/kleister/kleister-api)
[![Build Status](https://ci.appveyor.com/api/projects/status/vugtghebqi0xnfiw?svg=true)](https://ci.appveyor.com/project/kleisterz/kleister-api)
[![Stories in Ready](https://badge.waffle.io/kleister/kleister-api.svg?label=ready&title=Ready)](http://waffle.io/kleister/kleister-api)
[![Join the Matrix chat at https://matrix.to/#/#kleister:matrix.org](https://img.shields.io/badge/matrix-%23kleister-7bc9a4.svg)](https://matrix.to/#/#kleister:matrix.org)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/a7f12e6f17524e669df945546d4ee37c)](https://www.codacy.com/app/kleister/kleister-api?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=kleister/kleister-api&amp;utm_campaign=Badge_Grade)
[![Go Doc](https://godoc.org/github.com/kleister/kleister-api?status.svg)](http://godoc.org/github.com/kleister/kleister-api)
[![Go Report](http://goreportcard.com/badge/github.com/kleister/kleister-api)](http://goreportcard.com/report/github.com/kleister/kleister-api)
[![](https://images.microbadger.com/badges/image/kleister/kleister-api.svg)](http://microbadger.com/images/kleister/kleister-api "Get your own image badge on microbadger.com")

**This project is under heavy development, it's not in a working state yet!**

Kleister is a web UI to manage mod packs for the Minecraft, initially focused on the Technic Launcher and MCUpdater. Even if there is an upstream version available the Technic Launcher at [TechnicPack/TechnicSolder](https://github.com/TechnicPack/TechnicSolder) I prefered to implement it in Go for the API and VueJS for the UI including some further features like uploading the mods I want to manage and even generating docker images directly out of the managed packs. Hosting Minecraft servers based on docker images works pretty cool.

*Where does this name come from or what does it mean? It's quite simple, it's one german word for paste/glue, I thought it's a good match as it glues together the modpacks for Minecraft.*


## Install

You can download prebuilt binaries from the GitHub releases or from our [download site](http://dl.kleister.tech/api). You are a Mac user? Just take a look at our [homebrew formula](https://github.com/kleister/homebrew-kleister).


## Development

Make sure you have a working Go environment, for further reference or a guide take a look at the [install instructions](http://golang.org/doc/install.html). This project requires Go >= v1.8. It is possible to just execute `go get github.com/kleister/kleister-api/cmd/kleister-api`, but we prefer to use our `Makefile`:

```bash
go get -d github.com/kleister/kleister-api
cd $GOPATH/src/github.com/kleister/kleister-api

# install retool
make retool

# sync dependencies
make sync

# generate code
make generate

# build binary
make build

./bin/kleister-api -h
```


## Security

If you find a security issue please contact kleister@webhippie.de first.


## Contributing

Fork -> Patch -> Push -> Pull Request


## Authors

* [Thomas Boerger](https://github.com/tboerger)


## License

Apache-2.0


## Copyright

```
Copyright (c) 2018 Thomas Boerger <thomas@webhippie.de>
```

# Kleister: API server

[![Build Status](http://github.dronehippie.de/api/badges/kleister/kleister-api/status.svg)](http://github.dronehippie.de/kleister/kleister-api)
[![Go Doc](https://godoc.org/github.com/kleister/kleister-api?status.svg)](http://godoc.org/github.com/kleister/kleister-api)
[![Go Report](https://goreportcard.com/badge/github.com/kleister/kleister-api)](https://goreportcard.com/report/github.com/kleister/kleister-api)
[![Sourcegraph](https://sourcegraph.com/github.com/kleister/kleister-api/-/badge.svg)](https://sourcegraph.com/github.com/kleister/kleister-api?badge)
[![](https://images.microbadger.com/badges/image/kleister/kleister-api.svg)](http://microbadger.com/images/kleister/kleister-api "Get your own image badge on microbadger.com")
[![Join the chat at https://gitter.im/kleister/kleister](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/kleister/kleister)
[![Stories in Ready](https://badge.waffle.io/kleister/kleister-api.svg?label=ready&title=Ready)](http://waffle.io/kleister/kleister-api)

**This project is under heavy development, it's not in a working state yet!**

Where does this name come from or what does it mean? It's quite simple, it's one
german word for paste/glue, I thought it's a good match as it glues together the
modpacks for Minecraft.

Kleister is a web UI to manage mod packs for the Minecraft Technic launcher for
the Technic platform. Even if there is an upstream version available at
[TechnicPack/TechnicSolder](https://github.com/TechnicPack/TechnicSolder) I
prefered to implement it in Go for the API and with React for the UI including
some further features like uploading the mods I want to manage and even
generating docker images directly out of the managed packs. Hosting Minecraft
servers based on docker images works pretty cool.

The structure of the code base is heavily inspired by Drone, so those credits
are getting to [bradrydzewski](https://github.com/bradrydzewski), thank you for
this awesome project!


## Install

You can download prebuilt binaries from the GitHub releases or from our
[download site](http://dl.webhippie.de/kleister-api). You are a Mac user? Just take
a look at our [homebrew formula](https://github.com/kleister/homebrew-kleister).
If you are missing an architecture just write us on our nice
[Gitter](https://gitter.im/kleister/kleister-api) chat. Take a look at the help
output, you can enable auto updates to the binary to avoid bugs related to old
versions. If you find a security issue please contact thomas@webhippie.de first.


## Development

Make sure you have a working Go environment, for further reference or a guide
take a look at the [install instructions](http://golang.org/doc/install.html).
As this project relies on vendoring of the dependencies and we are not
exporting `GO15VENDOREXPERIMENT=1` within our makefile you have to use a Go
version `>= 1.6`. It is also possible to just simply execute the
`go get github.com/kleister/kleister-api/cmd/kleister-api` command, but we
prefer to use our `Makefile`:

```bash
go get -d github.com/kleister/kleister-api
cd $GOPATH/src/github.com/kleister/kleister-api
make clean build

./kleister-api -h
```


## Contributing

Fork -> Patch -> Push -> Pull Request


## Authors

* [Thomas Boerger](https://github.com/tboerger)


## License

Apache-2.0


## Copyright

```
Copyright (c) 2016 Thomas Boerger <http://www.webhippie.de>
```

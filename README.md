# Solder: API server

[![Build Status](http://github.dronehippie.de/api/badges/solderapp/solder-api/status.svg)](http://github.dronehippie.de/solderapp/solder-api)
[![Coverage Status](http://coverage.dronehippie.de//badges/solderapp/solder-api/coverage.svg)](https://aircover.co/solderapp/solder-api)
[![Go Doc](https://godoc.org/github.com/solderapp/solder-api?status.svg)](http://godoc.org/github.com/solderapp/solder-api)
[![Go Report](http://goreportcard.com/badge/solderapp/solder-api)](http://goreportcard.com/report/solderapp/solder-api)
[![Join the chat at https://gitter.im/solderapp/solder-api](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/solderapp/solder-api)
![Release Status](https://img.shields.io/badge/status-beta-yellow.svg?style=flat)

**This project is under heavy development, it's not in a working state yet!**

Solder is a web UI to manage mod packs for the Minecraft Technic launcher for
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

Make sure you have a working Go environment, for further reference or a guide
take a look at the [install instructions](http://golang.org/doc/install.html).
As this project relies on vendoring of the dependencies and we are not
exporting `GO15VENDOREXPERIMENT=1` within our makefile you have to use a Go
version `>= 1.6`

```bash
go get -d github.com/solderapp/solder-api
cd $GOPATH/src/github.com/solderapp/solder-api
make deps build

bin/drone-api -h
```

Later on we will also provide a download of prebuilt binaries for various
platforms, but this will start if we get to an somehow working state or if we
are more or less on feature parity with the upstream project.


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

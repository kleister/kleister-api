# Solder

[![Join the chat at https://gitter.im/solderapp/solder](https://badges.gitter.im/solderapp/solder.svg)](https://gitter.im/solderapp/solder?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

[![Build Status](http://github.dronehippie.de/api/badges/solderapp/solder/status.svg)](http://github.dronehippie.de/solderapp/solder)
[![Coverage Status](https://aircover.co/badges/solderapp/solder/coverage.svg)](https://aircover.co/solderapp/solder)
[![Join the chat at https://gitter.im/solderapp/solder](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/solderapp/solder)
![Release Status](https://img.shields.io/badge/status-beta-yellow.svg?style=flat)

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
As this project relies on vendoring of the dependencies you have to use a Go
version `>= 1.5`

```bash
go get github.com/solderapp/solder
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

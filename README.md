lddld
====

[![Build Status](https://travis-ci.org/Legenddigital/lddld.png?branch=master)](https://travis-ci.org/Legenddigital/lddld)
[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](http://copyfree.org)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/Legenddigital/lddld)

lddld is a Legenddigital full node implementation written in Go (golang).

This acts as a chain daemon for the [Legenddigital](https://Legenddigital.org) cryptocurrency.
lddld maintains the entire past transactional ledger of Legenddigital and allows
 relaying of transactions to other Legenddigital nodes across the world.  To read more
about Legenddigital please see the
[project documentation](https://docs.Legenddigital.org/#overview).

Note: To send or receive funds and join Proof-of-Stake mining, you will also 
need [lddlwallet](https://github.com/Legenddigital/lddlwallet).

This project is currently under active development and is in a Beta state.  It
is extremely stable and has been in production use since February 2016.

It is forked from [btcd](https://github.com/btcsuite/btcd) which is a bitcoin
full node implementation written in Go.  btcd is a ongoing project under active
development.  Because lddld is constantly synced with btcd codebase, it will
get the benefit of btcd's ongoing upgrades to peer and connection handling,
database optimization and other blockchain related technology improvements.

## Requirements

[Go](http://golang.org) 1.9 or newer.

## Getting Started

- lddld (and utilities) will now be installed in either ```$GOROOT/bin``` or
  ```$GOPATH/bin``` depending on your configuration.  If you did not already
  add the bin directory to your system path during Go installation, we
  recommend you do so now.

## Updating

#### Windows

Install a newer MSI

#### Linux/BSD/MacOSX/POSIX - Build from Source

- **Dep**

  Dep is used to manage project dependencies and provide reproducible builds.
  To install:

  `go get -u github.com/golang/dep/cmd/dep`

Unfortunately, the use of `dep` prevents a handy tool such as `go get` from
automatically downloading, building, and installing the source in a single
command.  Instead, the latest project and dependency sources must be first
obtained manually with `git` and `dep`, and then `go` is used to build and
install the project.

**Getting the source**:

For a first time installation, the project and dependency sources can be
obtained manually with `git` and `dep` (create directories as needed):

```
git clone https://github.com/Legenddigital/lddld $GOPATH/src/github.com/Legenddigital/lddld
cd $GOPATH/src/github.com/Legenddigital/lddld
dep ensure
go install . ./cmd/...
```

To update an existing source tree, pull the latest changes and install the
matching dependencies:

```
cd $GOPATH/src/github.com/Legenddigital/lddld
git pull
dep ensure
go install . ./cmd/...
```

For more information about Legenddigital and how to set up your software please go to
our docs page at
[docs.Legenddigital.org](https://docs.Legenddigital.org/getting-started/beginner-guide/).

## Docker

### Running lddld

You can run a Legenddigital node from inside a docker container.  To build the image
yourself, use the following command:

```
docker build -t Legenddigital/lddld .
```

Or you can create an alpine based image (requires Docker 17.05 or higher):

```
docker build -t Legenddigital/lddld:alpine -f Dockerfile.alpine .
```

You can then run the image using:

```
docker run Legenddigital/lddld
```

You may wish to use an external volume to customise your config and persist the
data in an external volume:

```
docker run --rm -v /home/user/lddldata:/root/.lddld/data Legenddigital/lddld
```

For a minimal image, you can use the Legenddigital/lddld:alpine tag.  This is typically
a more secure option while also being a much smaller image.

You can run lddlctl from inside the image.  For example, run an image (mounting
your data from externally) with:

```
docker run --rm -ti --name=lddld-1 -v /home/user/.lddld:/root/.lddld \
  Legenddigital/lddld:alpine
```

And then run lddlctl commands against it.  For example:

```
docker exec -ti lddld-1 lddlctl getbestblock
```


### Running Tests

All tests and linters may be run in a docker container using the script
`run_tests.sh`.  This script defaults to using the current supported version of
go.  You can run it with the major version of Go you would like to use as the
only arguement to test a previous on a previous version of Go (generally Legenddigital
supports the current version of Go and the previous one).

```
./run_tests.sh 1.9
```

To run the tests locally without docker:

```
./run_tests.sh local
```

## Contact

If you have any further questions you can find us at:

- irc.freenode.net (channel #Legenddigital)
- [webchat](https://webchat.freenode.net/?channels=Legenddigital)
- forum.Legenddigital.org
- Legenddigital.slack.com

## Issue Tracker

The [integrated github issue tracker](https://github.com/Legenddigital/lddld/issues)
is used for this project.

## Documentation

The documentation is a work-in-progress.  It is located in the
[docs](https://github.com/Legenddigital/lddld/tree/master/docs) folder.

## License

lddld is licensed under the [copyfree](http://copyfree.org) ISC License.

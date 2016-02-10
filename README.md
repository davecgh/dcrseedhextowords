dcrseedhextowords
=================

dcrseedhextowords is a simple utility that can convert a Decred seed from hex
to the seed words needed for importing into wallets.

## Requirements

[Go](http://golang.org) 1.5 or newer.

## Installation

#### Windows/Linux/MacOSX/POSIX - Zip Available

https://github.com/davecgh/dcrseedhextowords/releases

#### Build From Source

- Install Go according to the installation instructions here:
  http://golang.org/doc/install

- Ensure Go was installed properly and is a supported version:

```bash
$ go version
$ go env GOROOT GOPATH
```

NOTE: The `GOROOT` and `GOPATH` above must not be the same path.  It is
recommended that `GOPATH` is set to a directory in your home directory such as
`~/goprojects` to avoid write permission issues.

- Run the following command to obtain the utility, all dependencies, and install
  it:

```bash
$ go get -u github.com/davecgh/dcrseedhextowords/...
```

- dcrseedhextowords will now be installed in either ```$GOROOT/bin``` or
  ```$GOPATH/bin``` depending on your configuration.  If you did not already
  add the bin directory to your system path during Go installation, it is
  recommended to do so now.

## License

dcrseedhextowords is licensed under the [copyfree](http://copyfree.org) ISC
License.

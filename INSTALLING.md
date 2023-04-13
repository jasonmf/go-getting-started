# Installing Go

## Linux

Visit [the download page](https://go.dev/dl/) and download `go<version>.linux-<arch>.tar.gz` matching your machine. On an Intel machine, your `arch` is probably `amd64`. On a Raspberry Pi, it's probably `armv6l`.

If you have a previous version of Go installed, remove it:

```sh
sudo rm -rf /usr/local/go
```

Install the new version of Go:

```sh
sudo tar -C /usr/local xvf go<version>.linux-<arch>.tar.gz
```

If this is the first time installing Go on this machine, add `/usr/local/go/bin` to your `PATH` env var.

Verify the install:

```sh
go version
```

## Mac

Visit [the download page](https://go.dev/dl/) and download `go<version>.darwin-<arch>.pkg` matching your machine. On an M1/M2 machine, your `arch` is `arm64`. On an Intel machine, it's probably `amd64`.

Run the installer. If a version of Go is already installed, the new installer will cleanly replace it. No further configuration is required.

Verify the install:

```sh
go version
```

## Windows

Visit [the download page](https://go.dev/dl/) and download `go<version>.windows-amd64.msi`.

Run the installer. If a version of Go is already installed, the new installer will cleanly replace it. No further configuration is required.

Verify the install:

```sh
go version
```
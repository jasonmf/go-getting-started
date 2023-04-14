# Modules

Go code is organized into modules. A module is a folder with Go source code, a few special files, and maybe additional folders.

## Creating a module

First, you must decide on a module name. This can be changed later, though doing so can be inconvenient. If the code will be hosted on the Internet, like in a GitHub repo, it's common to name the module to match the URL where it will be hosted. This allows Go tools to easily find and download the module. A module is created by going into the folder where the code will live (or existing code lives) and running `go mod init <name>`.

This repo includes a `hello-world` module. It was created with:

```sh
$ go mod init github.com/jasonmf/go-getting-started/hello-world
go: creating new go.mod: module github.com/jasonmf/go-getting-started/hello-world
```

## go.mod and go.sum

The `go mod` command creates and maintains module files. There _must_ be a `go.mod` and there will often be a `go.sum`.

The `go.mod` file describes your module and its dependencies.

```sh
$ cat go.mod 
module github.com/jasonmf/go-getting-started/hello-world

go 1.20
```

This module assumes go version 1.20. If the way modules are managed changes between Go versions, this helps the tools know what to do. This module currently has no dependencies.

The `go.sum` module records cryptographic checksums for your modules dependencies. If your module has no dependencies, you may not have a `go.sum` file. The `hello-world` module has no dependencies and thus no `go.sum`. Recording these cryptographic checksums allows the `go` tool to verify that the code the module depends on hasn't been changed.

## Updating go.mod and go.sum

You rarely modify `go.mod` and `go.sum` by hand. When you add, remove, or change dependencies in your code, simply running `go mod tidy` will update these files for you, based on what your code is importing.

## Module versions

Go modules have versions and those versions are defined by a combination of repo tags (e.g. `git tag`) and folder names. The repo tags follow the [semver](https://semver.org/) standard.

A module with no version tags at all gets referenced by Go tools as `v0.0.0-<timestamp>-<checksum>`. These `v0.0.0` tags are generated automatically by `go mod tidy`.

As a matter of convention, modules with versions beginning with `v0.` (e.g. `v0.x.y`) are assumed to be subject to change; a newer `v0.x.y` may have changes that break existing code.

Similarly, modules beginning with `v1.` are expected to be _stable_; new features may be added and bugs fixed, but code that compiles using a `v1.` version should also compile against later `v1.` versions. Code using one major version of a module should always be able to compile against later versions of the module within the same major version.

Starting with `v2` and above, code is expected to live under a `vX` subfolder in addition to having the semver repo tag. For example, `v0.1.5` and `v1.3.6` of `github.com/foo/bar` can be found at `github.com/foo/bar`, as you would expect. Subsequent major versions must be in a subfolder matching that major version. So `v2.0.5` would be in `github.com/foo/bar/v2` and `v3.8.1` would be in `github.com/foo/bar/v3`. This was a pretty contraversial decision and it's reasonable to be annoyed by this.
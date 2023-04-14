# Compiling and Running Go Code

## Prerequisites

- You should have the source code for a Go program. You can use the code in [`hello-world`](hello-world/) in this repo.

## Running Go code

The easiest way to run Go code is with the `go run` command:

```sh
$ go run .
Hello world!
```

The `.` tells the Go command to consider all the source code in the current directory. If the program takes arguments, you can pass those arguments after the path (`.` in this case):

```sh
$ go run . -subj gopher
Hello gopher!
```

The `go run` tool assumes that arguments _before_ the path are for the Go tool itself:

```sh
$ go run -subj gopher .
flag provided but not defined: -subj
usage: go run [build flags] [-exec xprog] package [arguments...]
Run 'go help run' for details.
```

The `go run` command has no `-subj` flag, so it returns an error.

### What `go run` does

If you're accustomed to scripting languages like shell, python, or javascript, you might assume that `go run` interprets the source code; this is not the case. The `go run` command _compiles_ an executable from the source code, saves that to a temporary location, runs the executable, and then deletes the executable afterward. The Go compiler is so fast this whole process is quicker than the interpreters of other languages.

With that in mind, if you want to automate things with Go, you probably want to use `go run` during development but build and distribute an executable for release/production. It's more efficient and means you don't have to worry about getting the Go toolchain installed on other systems.

## Compiling Go code

To build a compiled program, use `go build`:

```sh
$ go build .
$ ls
go.mod          hello-world     main.go
```

By default, the compiled program will be named after the folder you're in. To choose a specific name, use the `-o <name>` flag:

```sh
$ go build -o hello .
$ ls
go.mod          hello           hello-world     main.go
```

### Static executables

By default, Go executables are _statically-compiled_, meaning they don't require any additional libraries on the system that runs them, aside from libraries that are part of the system itself. A Go-compiled executable can be run on other systems without needing to install dependencies on those other systems.

### Fully-static executables

Go executables can be compiled to not even require operating system libraries. This results in a larger executable but means the executable can even be run in a `FROM scratch` docker container. Except for the `FROM scratch` container case, you probably don't want to bother with this.

To do so:

```sh
$ CGO_ENABLED=0 go build -a -tags netgo -ldflags '-w -s' .
```

### Stripping debug symbols

By default, Go executables include information to help with debugging. This information can be omitted during the compilation process to save space:

```sh
$ go build -o with-debug .
$ go build -o without-debug -ldflags="-s -w" .
$ ls -lh with*
-rwxr-xr-x  1 jason  staff   2.0M Apr 14 13:33 with-debug
-rwxr-xr-x  1 jason  staff   1.5M Apr 14 13:33 without-debug
```

### Cross-compiling

Without any additional tooling, the `go build` command can build executables for other operating systems and architectures. This is controlled with the environment variables `GOOS` (for operating system) and `GOARCH` for architecture.

To build an executable for Windows on Intel:

```sh
$ GOOS=windows GOARCH=amd64 go build -o hello-win-amd64.exe .
$ file hello-win-amd64.exe 
hello-win-amd64.exe: PE32+ executable (console) x86-64 (stripped to external PDB), for MS Windows
```
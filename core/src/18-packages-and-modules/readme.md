![](/assets/gologo.png)

# The Go programming language - Packages & Modules

## Packages
A `package` is simply a directory containing `.go` files with the same package staement at the top line. A package contains one or many source files, source files must all have the `.go` extension of course and they must all beong to he same directory. A good package design is one that does one task, or avhieve one goal. This leads to concise understanding, greater maintainability and better re-use.

Packages come in two types:

- `executable` these belong to the `package main` and one of the source files must be `main.go` with a `main()` function.
- `non-executable` or `library` packages - these may have any name and are typically utility or goal achieving tasks. They cannot be executed standalone. Instead they are imported into your main program and used.

## The GOPATH
The `GOPATH` was once a requirement of installling go and although it may still be user defined there is an implicit `$HOME/go` path that is given to GOPATH if it is not user defined. TYicaly we have a 3-folder sub structure here with `pgk`, `bin` and `src`, the `src` is where Go developers usually keep there code. The release of Go1.11 made that a freer choice as prior ot that the location of go source code was mode prescribed than it is today.

The `go env` at the CLI will return all the envionment values applicable to the Go programming language for any system on which it is installed.

## Private and Public access

Casing is important in go in that capitalised functions are exported and are thereore reachable from outside.

## Import scope

The imports are file scope in Go.

## The init function

It is possible for an `init` function to be used in a Go prgram. It is even possible to have more than 1. Where mre than one is present they are read and executed top down fashion so file placement matters in this condition. Think of multiple inits in a file as a weakness and a code smell.

# Go Modules

Go modules allows for code organisation outside of the traditional GOPATH. This is a develper experience and project experience improvement, but modules also offer another key advantage, package versioning. In the GOPATH `go get` would retrieve, store and referto one version of a package of that name.

as of go 1.13 the go command downloads, checksums and verifies from the module mirror package versions. But, what is a module? It's a set of related go packages stored in a dorectory tree with `go.mod` and `go.sum` files at its root. A module contains _1-n_ go files and the go.mod defines the modules path, the import path and dependency requirements. Modules are directly suppported and resolve dependecies issues with other modules.

## Import modules

The go modules system allows us to locate our projects anywhere on our filesystem. However, other modules we impot will be located on the GOPATH in the `pkg/mod` directory.

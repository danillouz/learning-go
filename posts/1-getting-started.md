# Getting Started

> Sun 18 Feb 2018

Instructions are macOS specific.

## Install Go Tools

```
> brew install go
```

See [here](https://golang.org/doc/install) how to install on other Operating Systems.

## Setup Workspace

A workspace is a directory hierarchy with three directories at its root:

- `src`: source repositories
- `pkg`: package objects
- `bin`: executable commands

You should keep your Go source code and dependencies in a single workspace, which differs
from other programming environments in which every project has a separate workspace. The default
workspace, called `go`, is located in your home directory and is denoted by [\$GOPATH](https://golang.org/doc/code.html#GOPATH).

Create your workspace:

```
> mkdir -p $HOME/go/{src,pkg,bin}
```

More information about workspaces can be found [here](https://golang.org/doc/code.html#Workspaces).

## Add `$GOPATH` to `$PATH`

Add the following to `.zshrc`:

```
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

This allows you to executw Go binaries without specifying the full path to the binary. More information
can be found [here](https://github.com/golang/go/wiki/SettingGOPATH).

## Setup Code Editor

Install [Visual Studio Code Go Plugin](https://marketplace.visualstudio.com/items?itemName=lukehoban.Go).

More plugins can be found [here](https://golang.org/doc/editors.html).

## Code Organization

An _import path_ is a string that uniquely identifies a package. A package's import path corresponds
to its location inside a workspace or in a remote repository.

I keep my code on GitHub, therefore `github.com/danillouz` will be the _base path_ of all my packages.

## Hello World

First create a package called `hello`:

```
> mkdir -p $GOPATH/src/github.com/danillouz/hello
```

Create a file called `hello.go` inside the `hello` package with the following [contents](../code/hello/hello.go):

```go
package main

import "fmt"

func main() {
  fmt.Println("Hello World!")
}
```

Install the package; from inside the `hello` directory you can run:

```
> go install
```

Or from `$GOPATH/go/src` you can run:

```
> go install github.com/danillouz/hello
```

Now you can type the binary name `hello`:

```
> hello
Hello World!
```

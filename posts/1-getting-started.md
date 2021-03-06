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

## Running Tests

Test can be run in the `code/hello` dir by running:

```
> go test
```

This will run the tests as specified in `hello_test.go`.

In Go everything needed to run tests is built in to the language and works as
long as:

- the file to be tested ends with `_test.go`. For example `hello_test.go`.
- the test function must start with `Test`. For example `TestHello(t *testing.T)`.

Test functions are called with `t *testing.T` which is a type to help interact with the testing package.

## Documentation and Examples

Packages and code can be documented with comments in code. The tool `godoc` can
then be used to generate and serve documentation of all code in the workspace
locally by running:

```
> godoc -http :8080
```

Navigate to `localhost:8080/pkg` to view the docs.

Examples can be generated by defining a function that starts with `Example` inside
the `_test.go` file. For example `ExampleSum()`. Examples are tests, are compiled
and can optionally be exectued as part of the test suite. This happens when an
"Output" comment is present in the example:

```go
func ExampleSum() {
  sum := Sum(4, 4)
  fmt.Println(sum)
  // Output: 8
}
```

> When the example is executed the testing framework captures data written to standard output and then compares the output against the example's "Output:" comment. The test passes if the test's output matches its output comment.

When the "Output:" comment is removed, the code is only compiled, but not executed.

> Examples without output comments are useful for demonstrating code that cannot run as unit tests, such as that which accesses the network, while guaranteeing the example at least compiles.

More information can be found [here](https://blog.golang.org/examples).

## Benchmarks

Benchmarks are functions that start with the name `Benchmark` and will be executed
when `go test` is run with the `-bench` flag. For example:

```go
func BenchmarkSomeFunc(*b testing.B) {
  for i := 0; i < b.N; i++ {
    // Code to be benchmarked, e.g. SomeFunc()
  }
}
```

> The benchmark function must run the target code b.N times. During benchmark execution, b.N is adjusted until the benchmark function lasts long enough to be timed reliably.

Then run the benchmark with:

```
> go test -bench .
```

More information can be found [here](https://golang.org/pkg/testing/#hdr-Benchmarks).

## Test Coverage

To enable test coverage analytics run:

```
> go test -cover
```

## Race Detector

Go can detect race conditions with:

```
> go test -race
```

More information can be found [here](https://blog.golang.org/race-detector).

## Go Modules

Go modules allow you to organize Go code outside of `GOPATH`.

Initialize:

```
go mod init github.com/danillouz/spidey
```

Add deps to `go.mod`:

```
go get -u ./...
```

_Commit `go.mod` and `go.sum`._

For more information see [modules wiki](https://github.com/golang/go/wiki/Modules).

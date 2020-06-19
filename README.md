# Random

[![GoDoc](https://godoc.org/github.com/erdaltsksn/random?status.svg)](https://godoc.org/github.com/erdaltsksn/random)
![Go](https://github.com/erdaltsksn/random/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/erdaltsksn/random)](https://goreportcard.com/report/github.com/erdaltsksn/random)

Random is a library and cli app to generate random data. This app also known as
Faker.

## Features

- Cross-Platform
- Automatically copied into system clipboard

## Getting Started

### 1. Get the package and install the application.

```sh
go get github.com/erdaltsksn/random
```

## Installation

### Using Homebrew

```sh
brew install erdaltsksn/tap/random
```

### Using Go Modules

```sh
go get github.com/erdaltsksn/random
```

## Updating

### Using Homebrew

```sh
brew upgrade erdaltsksn/tap/random
```

### Using Go Modules

```sh
go get -u github.com/erdaltsksn/random
```

## Usage

### As a CLI App

You may find the documentation for [each command](docs/random.md) inside the
[docs](docs) folder.

### As a library

```sh
package main

import (
	"fmt"

	random "github.com/erdaltsksn/random/v1"
)

func main() {
    fmt.Println(random.Digit())
}
```

### Getting Help

```sh
random --help
random [command] --help
```

## Contributing

If you want to contribute to this project and make it better, your help is very
welcome. See [CONTRIBUTING](docs/CONTRIBUTING.md) for more information.

## Disclaimer

In no event shall we be liable to you or any third parties for any special,
punitive, incidental, indirect or consequential damages of any kind, or any
damages whatsoever, including, without limitation, those resulting from loss of
use, data or profits, and on any theory of liability, arising out of or in
connection with the use of this software.

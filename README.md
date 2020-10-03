# Random

[![PkgGoDev](https://pkg.go.dev/badge/github.com/erdaltsksn/random)](https://pkg.go.dev/github.com/erdaltsksn/random)
![Go](https://github.com/erdaltsksn/random/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/erdaltsksn/random)](https://goreportcard.com/report/github.com/erdaltsksn/random)
![CodeQL](https://github.com/erdaltsksn/random/workflows/CodeQL/badge.svg)

Random is a library and cli app to generate random data. This app also known as
Faker.

## Features

- Cross-Platform
- Automatically copied into system clipboard

## Getting Started

### 1. Install the application.

```sh
brew install erdaltsksn/tap/random
```

### 2. Run the CLI application.

```sh
random <COMMAND>
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

## Updating / Upgrading

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

	"github.com/erdaltsksn/random"
)

func main() {
    fmt.Println(random.Digit())
}
```

## Getting Help

```sh
# Getting help for related command.
random --help
random [command] --help

# Show available `make` commands.
make help
```

## Contributing

If you want to contribute to this project and make it better, your help is very
welcome. See [CONTRIBUTING](.github/CONTRIBUTING.md) for more information.

## Security Policy

If you discover a security vulnerability within this project, please follow our
[Security Policy Guide](.github/SECURITY.md).

## Code of Conduct

This project adheres to the Contributor Covenant [Code of Conduct](.github/CODE_OF_CONDUCT.md).
By participating, you are expected to uphold this code.

## Disclaimer

In no event shall we be liable to you or any third parties for any special,
punitive, incidental, indirect or consequential damages of any kind, or any
damages whatsoever, including, without limitation, those resulting from loss of
use, data or profits, and on any theory of liability, arising out of or in
connection with the use of this software.

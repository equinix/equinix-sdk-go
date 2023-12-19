# Equinix Go SDK

[![Maintained](https://img.shields.io/badge/stability-maintained-green.svg)](https://github.com/equinix-labs/equinix-labs/blob/main/maintained-statement.md)
[![Release](https://img.shields.io/github/v/release/equinix/equinix-sdk-go)](https://github.com/equinix/equinix-sdk-go/releases/latest)
[![GoDoc](https://godoc.org/github.com/equinix/equinix-sdk-go?status.svg)](https://godoc.org/github.com/equinix/equinix-sdk-go)

This is the official Go SDK for Equinix services.  This SDK is currently provided with a major version of [v0](https://blog.golang.org/v2-go-modules). We aim to avoid breaking changes to this library, but they will certainly happen as we work towards a stable v1 library.

Each Equinix service supported by this SDK is maintained as a separate package that is generated from the OpenAPI specification for that service.  If any Equinix service is not supported by this SDK and you would like to see it added, please [submit a change request](CONTRIBUTING.md)

## Installation

To import this library into your Go project:

```go
import "github.com/equinix/equinix-sdk-go"
```

Download the module with:

```sh
go get github.com/equinix/equinix-sdk-go
```

## Usage

You can see usage of the generated code in the [`examples` directory](https://github.com/equinix/equinix-sdk-go/tree/main/examples).

Full package documentation is available at <https://pkg.go.dev/github.com/equinix/equinix-sdk-go>.

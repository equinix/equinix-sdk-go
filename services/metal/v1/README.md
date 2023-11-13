# Golang client for Equinix Metal

[![Experimental](https://img.shields.io/badge/Stability-Experimental-red.svg)](https://github.com/equinix-labs/equinix-labs/blob/main/uniform-standards.md)

> **[Experimental](https://github.com/equinix-labs/equinix-labs/blob/main/experimental-statement.md)**
> This is experimental. Don't use it in production. Examples demonstrate that this client is usable. Please submit patches and open issues with your experience. This repo contains Go code generated from a [customized OpenAPI specification](spec/oas3.patched) based on the [Equinix Metal API spec](https://api.equinix.com/metal/v1/api-docs).  The client is generated using the Go client support built into the [OpenAPITools openapi-generator](https://github.com/OpenAPITools/openapi-generator).

## Contents

- `Makefile` includes tasks to fetch the API spec, apply patches, and generate a client
- `spec/oas3.fetched` a directory of the latest fetched OpenAPI spec
- `spec/oas3.patched` a directory of the latest patched OpenAPI spec
- `patches/spec.fetched.json/*.patch` patch files to apply against the fetched OpenAPI spec
- `patches/post/*.patch` patch files to apply against the generated Go code
- `examples/` hand crafted examples to demonstrate usage
- `metal/v1/` generated Go client

## Generated Client

See [API.md](API.md)

## Build

To build the client, run `make`.

## Examples

You can see usage of the generated code in the [`examples` directory](https://github.com/equinix-labs/metal-go/tree/main/examples). In order to try, export `METAL_AUTH_TOKEN` token and execute the code, e.g.

## Patches

1. Make changes in ``spec/oas3.patched/`` dir.
2. Create a patch file in metal-go:
   ```
   git diff spec/oas3.patched/ > ../patches/spec.fetched.json/<patchfilename>
   cd ..
   ```
3. ``patchfilename`` should be in format: ``<patch_index>-<short_patch_decription_or_identifier>.patch``
4. Run ``make`` to reapply the changes to the fetched OpenAPI spec and confirm that the patched spec includes the expected changes



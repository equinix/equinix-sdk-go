# GitHub Copilot Instructions

## Project Overview

This is `equinix-sdk-go`, the official Go SDK for Equinix services. The SDK is **generated** from OpenAPI specifications — most Go source files under `services/` are produced by `openapi-generator` and should not be edited directly.

## Architecture

- Each Equinix service is a separate Go package under `services/<service>/`
- OpenAPI specs live under `spec/services/<service>/`
- Code generation is driven by `Makefile.<service>` files and the shared `Makefile`
- Docker is used to run `openapi-generator` and other tooling

## Coding Conventions

- **Do not manually edit** files under `services/<service>/` — they are generated and will be overwritten
- Service names are lowercase alphanumeric (e.g. `metalv1`, `fabricv4`, `networkedgev1`)
- Use [Conventional Commits](https://www.conventionalcommits.org/): `feat:`, `fix:`, `sync:`, `chore:`, `docs:`
- Go code must pass `golangci-lint` — run `make lint` to check
- Tests are run with `make test`

## Common Workflows

### Add a new Equinix service to the SDK

Use the GitHub Actions workflow **Onboard new API** or run:
```bash
make onboard-service
```
Provide a lowercase service name and the URL to the OpenAPI spec.
OpenAPI specs for Equinix services follow the pattern:
```
https://docs.equinix.com/api-catalog/{service}/openapi.yaml
```

### Regenerate code after spec changes

```bash
make -f Makefile.<service> generate
```

### Patch a spec

Create patch files under `spec/services/<service>/patches/` and run:
```bash
make -f Makefile.<service> patch
```

## Key Files

| File | Purpose |
|------|---------|
| `Makefile` | Top-level tasks |
| `Makefile.<service>` | Per-service tasks |
| `templates/Makefile.sdk` | Template for new service Makefiles |
| `templates/.github/workflows/sync.yaml` | Template for new sync workflows |
| `script/onboard_service.sh` | Interactive script to onboard a service |
| `config/openapi-generator.json` | openapi-generator configuration |
| `version` | Current SDK version |
| `AGENTS.md` | AI agents guide |
| `.github/agents/onboard-api.md` | Agent skill for onboarding APIs |

## What Copilot Should Help With

- Writing spec patches (`.patch` files) for OpenAPI spec issues
- Writing custom `openapi-generator` templates under `templates/services/<service>/`
- Writing hand-crafted usage examples under `examples/services/<service>/`
- Writing post-generation patches under `patches/services/<service>/`
- Updating `Makefile` targets or `Makefile.<service>` configurations
- Writing GitHub Actions workflows following the patterns in `.github/workflows/`

## What Copilot Should Avoid

- Directly modifying generated Go files under `services/`
- Introducing new Go dependencies without updating `go.mod` via `make mod`
- Suggesting edits to `go.sum` directly

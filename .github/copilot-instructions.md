# Equinix Go SDK - Copilot Instructions

## Repository Overview

Official Go SDK for Equinix services. Generates Go client libraries from OpenAPI specs for 10 services (metalv1, fabricv4, eiav2, accesstokenv1, lookupv2, orderhistoryv1, ordersv2, securecabinetv1, smarthandsv1, stsv1alpha). **~1,375 generated Go files**. Minimal hand-written code.

- **Language:** Go 1.19 | **Codegen:** OpenAPI Generator v7.12.0 | **Runtime:** Docker | **Version:** 0.61.0
- **Structure:** Multi-service SDK with per-service Makefiles (Makefile.<service>)

## Project Structure

**Key Directories:**
- `Makefile` + `Makefile.<service>` - Build orchestration (10 per-service configs)
- `services/<service>/` - Generated Go clients (**DO NOT manually edit**)
- `spec/services/<service>/` - OpenAPI specs (oas3.fetched/, oas3.patched/, patches/)
- `templates/services/<service>/` - Custom openapi-generator templates
- `examples/services/<service>/` - Hand-written usage examples
- `extensions/equinixoauth2/` - Hand-written OAuth2 utilities
- `config/openapi-generator.json`, `.golangci.yaml` - Tool configurations
- `.github/workflows/` - CI/CD (test-codegen, test-examples, validate_pr, sync-*, release)

## Build Commands

**ALL commands run in Docker containers.** First-time image pulls add 1-3 minutes.

### Standard Workflow (in order):

1. **Dependencies** - `make mod` (~30s) - ALWAYS run first after checkout
2. **Verify Patches** - `make patch-all` (~10s) - Validates patch files apply cleanly
3. **Generate Code** - `make generate-all` (**5-10 min all services**) OR `make -f Makefile.<service> generate` (~30s single)
4. **Format** - `make fmt` (~20s) - Runs gofumpt v0.3.1
5. **Test** - `make test` (~10s) - Currently no test files
6. **Lint** - `make lint` (**40-60s**) - golangci-lint v2.3.0, must show "0 issues"
7. **Build Examples** - `make build-examples-all` (~60s)

### Per-Service Targets:
- `make -f Makefile.<service> fetch|patch|generate|validate|build-examples`

**CRITICAL:** `make generate-all` must produce zero git diffs (CI requirement). metalv1 uses special two-step generation (v7.4.0 for merge, v7.12.0 for codegen) due to swagger-parser bug.

## CI/CD Workflows

### Every Push/PR:
1. **test-codegen.yaml** - Main gate: patch-all → generate-all → **fails if uncommitted changes** → lint
2. **test-examples.yaml** - Validates all examples build
3. **validate_pr.yaml** - **PR title MUST follow Conventional Commits** (feat:, fix:, docs:, chore:)

### Manual (workflow_dispatch):
- **sync-<service>.yaml** - Updates service from latest OpenAPI spec
- **release.yaml** - Creates new release with Semantic Release

## Common Pitfalls & Solutions

### Generated Code Changes
**Problem:** Most code in `services/` is **generated**. Manual edits are overwritten.  
**Solution:** 
- Modify `spec/services/<service>/patches/*.patch` (OpenAPI spec patches)
- OR add custom templates in `templates/services/<service>/`
- Then: `make -f Makefile.<service> generate`

### Creating Patches
1. Edit `spec/services/<service>/oas3.patched/` files
2. Create patch: `git diff spec/services/<service>/oas3.patched > spec/services/<service>/patches/<index>-<desc>.patch`
3. Naming: `<index>-<description>.patch` (e.g., `20240119-paginate-ips.patch`)
4. Verify: `make -f Makefile.<service> patch` then `make -f Makefile.<service> generate`

### "Uncommitted Changes After Generation" CI Error
**Cause:** Generated code out of sync with specs/templates.  
**Fix:** Run `make generate-all` locally, commit ALL changes including generated code. Never manually edit to fix.

### Lint Failures
Run `make lint` locally (60s). For hand-written code, run `make fmt` first. Generated code is pre-compliant.

### Patch Application Failures
Check `make patch-all` for offset/fuzz warnings. Patches may need regeneration from updated specs.

### Docker Issues
- Ensure Docker daemon running, user has permissions (runs as UID:GID 1001:1001)
- Custom CA certs: Set `DOCKER_EXTRA_ARGS` (see DEVELOPMENT.md)

### Examples
- Binaries gitignored automatically
- Each example in `examples/services/<service>/<name>/` has standalone `main.go` + `go.mod`
- Build: `cd examples/services/<service>/<name> && go build .`

## Dependencies & Tools

**Go Dependencies (go.mod):** `golang.org/x/oauth2 v0.26.0`, `gopkg.in/validator.v2 v2.0.1`

**Docker Images (auto-pulled):**
- `openapitools/openapi-generator-cli:v7.12.0` (v7.4.0 for metalv1 merge)
- `golang:1.19`, `golangci/golangci-lint:v2.3.0`, `mikefarah/yq:4.30.8`

**Linters (.golangci.yaml):** errcheck, govet, ineffassign, staticcheck, unused, goimports

## Service Onboarding
`make onboard-service` → provide service name (lowercase) + OpenAPI spec URL → auto-creates Makefile.<service>, directories, sync workflow

## Pre-Submission Checklist
- [ ] `make mod` (if dependencies changed)
- [ ] `make patch-all` (patches apply cleanly)
- [ ] `make generate-all` (if specs/templates changed)
- [ ] `make lint` (0 issues required)
- [ ] `make build-examples-all` (compilation check)
- [ ] `git status --porcelain` empty (no uncommitted generated code)
- [ ] PR title follows Conventional Commits (feat:|fix:|docs:|chore:)
- [ ] `.gitignore` excludes artifacts (example binaries, .certs, temp files)

## Trust These Instructions
Comprehensive and tested. Search only if incomplete, errors occur, or repo structure changed significantly. Check `Makefile.<service>` or `.github/workflows/test-codegen.yaml` for specifics.

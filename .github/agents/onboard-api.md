# Skill: Onboard a New Equinix API Service

## Description

Onboard a new Equinix API service into the `equinix-sdk-go` SDK by scaffolding the required files and generating the initial Go client from an OpenAPI specification.

## When to Use

Use this skill when:
- A new Equinix API has been published and should be added to the SDK
- An OpenAPI spec is available at a public URL (no authentication required)
- You have the service name and the URL to its OpenAPI specification

The OpenAPI spec URL for Equinix services typically follows the pattern:
```
https://docs.equinix.com/api-catalog/{service}/openapi.yaml
```

## Inputs

| Input | Description | Example |
|-------|-------------|---------|
| `service_name` | Lowercase alphanumeric name for the service. Base it on the API path (e.g. `mysvcv1` for `api.equinix.com/mysvc/v1`). | `networkedgev1` |
| `spec_url` | Full URL to the root OpenAPI spec file | `https://docs.equinix.com/api-catalog/networkedge/openapi.yaml` |

## Steps

1. **Validate inputs**
   - Confirm `service_name` is lowercase alphanumeric (no spaces or special characters)
   - Confirm `Makefile.<service_name>` does not already exist (service not already onboarded)
   - Confirm `spec_url` points to a publicly accessible OpenAPI spec

2. **Scaffold Makefile**
   - Copy `templates/Makefile.sdk` → `Makefile.<service_name>`
   - Replace `__PACKAGE_NAME__` with the service name
   - Replace `__SPEC_BASE_URL__` with the base URL (everything before the filename)
   - Replace `__SPEC_ROOT_FILE__` with the spec filename (e.g. `openapi.yaml`)

3. **Scaffold sync GitHub Actions workflow**
   - Copy `templates/.github/workflows/sync.yaml` → `.github/workflows/sync-<service_name>.yaml`
   - Replace `__PACKAGE_NAME__` with the service name

4. **Create required directories**
   ```
   spec/services/<service_name>/          (with .keep placeholder)
   templates/services/<service_name>/     (with .keep placeholder)
   patches/services/<service_name>/       (with .keep placeholder)
   ```

5. **Fetch the OpenAPI spec**
   ```bash
   make -f Makefile.<service_name> fetch
   ```

6. **Apply patches** (none initially, but run to verify)
   ```bash
   make -f Makefile.<service_name> patch
   ```

7. **Generate the Go client**
   ```bash
   make -f Makefile.<service_name> generate
   ```

8. **Stage changes for review**
   ```bash
   git add --intent-to-add .
   ```

## Automated Workflow

This skill is automated as a GitHub Actions workflow. Trigger it at:

**Actions → Onboard new API → Run workflow**

Or use the GitHub CLI:
```bash
gh workflow run onboard-api.yaml \
  -f service_name=<service_name> \
  -f spec_url=https://docs.equinix.com/api-catalog/<service>/openapi.yaml
```

## Output

Upon success, the following files are created or modified:

| File | Description |
|------|-------------|
| `Makefile.<service_name>` | Per-service Make tasks |
| `.github/workflows/sync-<service_name>.yaml` | Sync workflow to keep the client up to date |
| `spec/services/<service_name>/oas3.fetched/` | Downloaded OpenAPI spec |
| `spec/services/<service_name>/oas3.patched/` | Patched spec (initially identical to fetched) |
| `spec/services/<service_name>/patches/` | Empty patches directory |
| `templates/services/<service_name>/` | Empty custom templates directory |
| `patches/services/<service_name>/` | Empty post-generation patches directory |
| `services/<service_name>/` | Generated Go client code |

## Troubleshooting

- **Fetch fails**: Confirm the spec URL is publicly accessible and returns a valid OpenAPI YAML/JSON file.
- **Generate fails**: The spec may have issues incompatible with `openapi-generator`. Consider adding a spec patch in `spec/services/<service_name>/patches/`.
- **Patch fails**: Review existing patches — the spec may have changed in a way that makes a patch no longer apply cleanly.

## References

- [DEVELOPMENT.md](../../DEVELOPMENT.md) — Full developer guide
- [AGENTS.md](../../AGENTS.md) — AI agents guide for this repository
- [openapi-generator docs](https://openapi-generator.tech/docs/generators/go)
- [Equinix API Catalog](https://docs.equinix.com/api-catalog/)

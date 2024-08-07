# Development

## Repository Contents

- `Makefile` includes tasks to generate and validate the full SDK and to onboard a new service to the SDK
- `Makefile.<service>` includes tasks for generating and validating a single service from its OpenAPI spec
- `spec/services/<service>/oas3.fetched` a directory of the latest fetched OpenAPI spec for a single service
- `spec/services/<service>/oas3.patched` a directory of the latest patched OpenAPI spec for a single service
- `spec/services/<service>/patches/*.patch` patch files to apply against the fetched OpenAPI spec for a single service
- `examples/` hand crafted examples to demonstrate usage
- `services/<service>` generated Go client for a single service

## Optional: customize how Makefile runs Docker containers

If need to change how the `make` tasks run the containers used in this repo, you can use the `DOCKER_EXTRA_ARGS` environment variable that the Make tasks will pass to the `docker` CLI.

For example, if you have custom CA certificates on your host machine that you want to use inside the containers, you can run the following commands before executing any `make` tasks:

### macOS

```bash
# Make a certificate directory that is not inside any git repo
export CUSTOM_CERT_DIR="${HOME}/.certs"
export CA_FILE_NAME="ca-certificates.crt"
mkdir -p ${CUSTOM_CERT_DIR}

# Export CA certificates from system keychain into certificate directory
security export -t certs -f pemseq -k /System/Library/Keychains/SystemRootCertificates.keychain -o ${CUSTOM_CERT_DIR}/${CA_FILE_NAME}
# Append CA certificates from user keychain into CA file created above
security export -t certs -f pemseq -k /Library/Keychains/System.keychain >> $CUSTOM_CERT_DIR/${CA_FILE_NAME}

export DOCKER_EXTRA_ARGS="-v $(realpath ${CUSTOM_CERT_DIR}):/etc/ssl/certs"

# Run `make` tasks as documented below
```

### Windows

TBD

### Linux

TBD

## Build

To build the full SDK, run `make generate-all`.

To build the client for a single service, run `make -f Makefile.<service> generate`

## Onboard a new service

To add support for a new Equinix service to this SDK, you first need to publish a valid, complete OpenAPI spec for that service.  Once the OpenAPI spec is published:

1. Run `make onboard-service`
2. Specify a name for the service
    - The name should be lower-case alphanumeric; we recommend basing the name on the API URL.  For example, given a service `api.equinix.com/mysvc/v9`, a good service name would be `mysvcv9`
3. Specify the URL to the OpenAPI spec for the service
    - The OpenAPI spec must be served on a URL that is publicly available and does not require authentication

The onboarding process will create a new file `Makefile.<service>`, where `<service>` is the name you specified in (2) above.

The onboarding proces will automatically attempt to download the spec and generate the initial client. If initial generation fails, you may need to add custom templates or patch the OpenAPI spec for your service.

## Adding custom openapi-generator templates for a service

This repo uses [`openapi-generator`](https://github.com/OpenAPITools/openapi-generator) to generate Go code.  When necessary, we can provide custom templates to the generator in order to change the generated code.

When a service is onboarded, an empty `templates/services/<service>/` directory is created. To modify generated code, you can put a custom template in this directory.  Templates that are known to the generator can be seen here: https://github.com/OpenAPITools/openapi-generator/tree/master/modules/openapi-generator/src/main/resources/go

Before adding custom templates for a service, review [the `openapi-generator` docs for templating](https://openapi-generator.tech/docs/templating).

## Patch the OpenAPI spec for a service

While we aim to publish OpenAPI specs that can be used as-is for code generation, we sometimes need to make changes to a service's spec within the SDK in order to address language-specific concerns or work around an issue while waiting for a long-term fix.  When necessary, you can patch the spec for a single service as follows:

1. Make changes in ``spec/services/<service>/oas3.patched/`` dir.
2. Create a patch file from the modified spec dir:
   ```
   git diff spec/services/<service>/oas3.patched > spec/services/<service>/patches/<patchfilename>
   cd ..
   ```
3. ``patchfilename`` should be in format: ``<patch_index>-<short_patch_decription_or_identifier>.patch``
4. Run ``make -f Makefile.<service> patch`` to reapply the changes to the fetched OpenAPI spec and confirm that the patched spec includes the expected changes
5. Run `make -f Makefile.<service> generate` to regenerate the code using the patched spec and confirm that the code generates successfully and works as expected

## Update the OpenAPI spec for a service

The steps below may vary by service, depending on decisions made by the CODEOWNERS for that service, but by default the process for pulling in the latest OpenAPI specification for a service and regenerating the code is:

1. Publish the updated OpenAPI specification (that happens outside of this repository)
2. If necessary, update the `SPEC_BASE_URL` and `SPEC_ROOT_FILE` to match the URL where the updated OpenAPI specification is published
3. Go to the [Actions page](https://github.com/equinix/equinix-sdk-go/actions) and find the sync workflow for your service, and click the workflow name.  The name will be in the format `Sync <service> API spec`.
4. Click the gray `Run workflow` button, check that `main` is selected in the popup, and then click the green `Run workflow` button in the popup to run the sync workflow
5.  The sync workflow will open a Pull Request (PR) that includes the OpenAPI specification updates as well as any code changes caused by the update.

After merging your service changes you can release a new version of this SDK that includes your updates.

## Release the SDK

This SDK uses a Semantic Release action to automatically determine the next release version based on Conventional Commit tags in commit messages.  There is a push-button workflow for releasing the SDK after changes are made:

1. Navigate to the [Release workflow](https://github.com/equinix/equinix-sdk-go/actions/workflows/release.yaml).
2. Click the gray `Run workflow` button, check that `main` is selected in the popup, and then click the green `Run workflow` button in the popup to run the sync workflow

The release workflow will determine the next version, regenerate the code with the updated version number, and create a [GitHub Release](https://github.com/equinix/equinix-sdk-go/releases).

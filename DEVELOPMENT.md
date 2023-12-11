# Development

## Contents

- `Makefile` includes tasks to generate and validate the full SDK and to onboard a new service to the SDK
- `Makefile.<service>` includes tasks for generating and validating a single service from its OpenAPI spec
- `spec/services/<service>/oas3.fetched` a directory of the latest fetched OpenAPI spec for a single service
- `spec/services/<service>/oas3.patched` a directory of the latest patched OpenAPI spec for a single service
- `spec/services/<service>/patches/*.patch` patch files to apply against the fetched OpenAPI spec for a single service
- `examples/` hand crafted examples to demonstrate usage
- `services/<service>` generated Go client for a single service

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

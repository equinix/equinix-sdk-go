#!/bin/bash
#
# Lists APIs that appear in the Equinix API catalog (RFC 9727) but are not yet
# onboarded to this SDK.
#
# Output: a compact JSON array on stdout, suitable for a GitHub Actions matrix:
#   [{"name":"foov1","spec_url":"https://docs.equinix.com/api-catalog/foov1/openapi.yaml"}, ...]
#
# An API's name is the last segment of its catalog anchor
# (https://docs.equinix.com/api-catalog/<name>/). An API is "in the SDK" when
# Makefile.<name> exists in the repo root (this is what script/onboard_service.sh creates).
#
# Environment:
#   CATALOG_URL  catalog location (default: https://docs.equinix.com/.well-known/api-catalog)
#   CATALOG_FILE read the catalog from this local file instead of fetching it (for testing)
#   GH_TOKEN     if set, APIs that already have an open onboarding PR (branch onboard/<name>) are skipped
#
# Skipped APIs are reported on stderr as GitHub Actions annotations.

set -euo pipefail

CATALOG_URL="${CATALOG_URL:-https://docs.equinix.com/.well-known/api-catalog}"
ANCHOR_PREFIX="https://docs.equinix.com/api-catalog/"

if [ -n "${CATALOG_FILE:-}" ]; then
  catalog=$(cat "$CATALOG_FILE")
else
  catalog=$(curl --fail --silent --show-error --location --max-time 60 "$CATALOG_URL")
fi

open_branches=""
if [ -n "${GH_TOKEN:-}" ]; then
  open_branches=$(gh pr list --state open --limit 500 --json headRefName --jq '.[].headRefName')
fi

# One tab-separated "anchor<TAB>spec_url" line per catalog entry that has an OpenAPI description.
entries=$(jq -r '
  .linkset[]
  | select(.anchor != null and ((.["service-desc"] // []) | length) > 0)
  | [.anchor, .["service-desc"][0].href]
  | @tsv' <<<"$catalog")

result="[]"
while IFS=$'\t' read -r anchor spec_url; do
  [ -n "$anchor" ] || continue

  if [[ "$anchor" != "$ANCHOR_PREFIX"* ]]; then
    echo "::warning::Skipping catalog anchor with unexpected format: ${anchor}" >&2
    continue
  fi
  name="${anchor%/}"
  name="${name##*/}"

  # The name becomes a Go package name, a Makefile suffix, and a directory name.
  if [[ ! "$name" =~ ^[a-z0-9]+$ ]]; then
    echo "::warning::Skipping '${name}': not a valid SDK service name (must be lower-case alphanumeric)" >&2
    continue
  fi
  if [[ "$spec_url" != "$ANCHOR_PREFIX"* ]]; then
    echo "::warning::Skipping '${name}': unexpected spec URL ${spec_url}" >&2
    continue
  fi
  # Some catalog entries describe event-driven APIs (AsyncAPI), which openapi-generator can't consume.
  if [[ "${spec_url##*/}" == asyncapi.* ]]; then
    echo "::notice::Skipping '${name}': spec is AsyncAPI, not OpenAPI (${spec_url})" >&2
    continue
  fi

  if [ -f "Makefile.${name}" ]; then
    continue
  fi
  if grep -qxF "onboard/${name}" <<<"$open_branches"; then
    echo "::notice::Skipping '${name}': an onboarding PR is already open" >&2
    continue
  fi

  result=$(jq -c --arg name "$name" --arg spec_url "$spec_url" \
    '. + [{name: $name, spec_url: $spec_url}]' <<<"$result")
done <<<"$entries"

echo "$result"

#!/bin/bash
#
# Compares the Equinix API catalog (RFC 9727) with the APIs in this SDK and prints
# the work to do, as a compact JSON array for a GitHub Actions matrix:
#   [{"name":"foov1","action":"onboard","spec_url":"https://...","branch":"onboard/foov1","title":"feat: ..."}, ...]
#
# An API's name is the last segment of its catalog anchor
# (https://docs.equinix.com/api-catalog/<name>/). An API is "in the SDK" when
# Makefile.<name> exists in the repo root (script/onboard_service.sh creates it).
#
#   action "onboard"  the API is in the catalog but not in the SDK
#   action "sync"     the API is in both: refresh its spec and regenerate the code
#
# APIs that are in the SDK but not in the catalog are only reported. They are never removed.
#
# Environment:
#   CATALOG_URL  catalog location (default: https://docs.equinix.com/.well-known/api-catalog)
#   CATALOG_FILE read the catalog from this local file instead of fetching it (for testing)
#   GH_TOKEN     if set, the script checks for open PRs on each API's branch. An API with an
#                open PR stays in the plan, so that the job refreshes the PR with the current
#                catalog spec. The job changes nothing when the PR is already up to date.
#                The script skips an API when a person added commits to its PR branch,
#                so that the job does not overwrite manual fixes.
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

# One "branch<TAB>number" line for each open PR.
open_prs=""
if [ -n "${GH_TOKEN:-}" ]; then
  open_prs=$(gh pr list --state open --limit 500 --json headRefName,number --jq '.[] | "\(.headRefName)\t\(.number)"')
fi

# Succeeds when someone other than the workflow bot authored a commit on the PR.
# A commit author with no GitHub login counts as a person.
pr_has_manual_commits() {
  gh pr view "$1" --json commits --jq '.commits[].authors[].login // ""' \
    | grep -qvxF 'github-actions[bot]'
}

# One tab-separated "anchor<TAB>spec_url" line per catalog entry that has an API description.
entries=$(jq -r '
  .linkset[]
  | select(.anchor != null and ((.["service-desc"] // []) | length) > 0)
  | [.anchor, .["service-desc"][0].href]
  | @tsv' <<<"$catalog")

result="[]"
catalog_names=" "
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

  catalog_names+="${name} "

  if [ -f "Makefile.${name}" ]; then
    action="sync"
    branch="catalog-sync/${name}"
    title="feat: sync ${name} API with the API catalog"

    # The sync uses the spec URL in Makefile.<name>. Report when it differs from the catalog.
    makefile_url="$(sed -n 's/^SPEC_BASE_URL:=\(.*\)/\1/p' "Makefile.${name}")/$(sed -n 's/^SPEC_ROOT_FILE:=\(.*\)/\1/p' "Makefile.${name}")"
    if [ "$makefile_url" != "$spec_url" ]; then
      echo "::warning::'${name}': Makefile spec URL (${makefile_url}) differs from the catalog (${spec_url}). Syncing from the Makefile URL." >&2
    fi
  else
    action="onboard"
    branch="onboard/${name}"
    title="feat: onboard ${name} API"
  fi

  pr_number=$(awk -F'\t' -v b="$branch" '$1 == b { print $2; exit }' <<<"$open_prs")
  if [ -n "$pr_number" ]; then
    if pr_has_manual_commits "$pr_number"; then
      echo "::notice::Skipping '${name}': PR #${pr_number} on branch ${branch} has manual commits" >&2
      continue
    fi
    echo "::notice::'${name}': PR #${pr_number} is open on branch ${branch}. The job refreshes it." >&2
  fi

  result=$(jq -c --arg name "$name" --arg action "$action" --arg spec_url "$spec_url" \
    --arg branch "$branch" --arg title "$title" \
    '. + [{name: $name, action: $action, spec_url: $spec_url, branch: $branch, title: $title}]' <<<"$result")
done <<<"$entries"

# Report SDK APIs that the catalog no longer lists. Nothing is removed automatically.
for makefile in Makefile.*; do
  name="${makefile#Makefile.}"
  if [[ "$catalog_names" != *" ${name} "* ]]; then
    echo "::notice::'${name}' is in the SDK but not in the API catalog. It was not synced or removed." >&2
  fi
done

echo "$result"

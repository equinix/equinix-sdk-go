#!/bin/bash
# Onboard a new service to the SDK.
#
# Usage: script/onboard_service.sh [<service> [<spec-url>]]
#
#   service   short lowercase name, used as the Go package name (for example fabricv4)
#   spec-url  URL of the OpenAPI spec. The script asks for any argument that is missing.
#
# Set SKIP_GENERATE=1 to create the scaffolding only. The script then does not fetch
# the spec or generate code. Run "make -f Makefile.<service> fetch patch generate" afterwards.

set -eu


function replace_in_file() {
  local file="$1"
  local search="$2"
  local replace="$3"

  # note that the following approach works on both GNU and BSD/Mac sed:
  # see https://stackoverflow.com/questions/5694228/sed-in-place-flag-that-works-both-on-mac-bsd-and-linux
  sed -i.bak "s|${search}|${replace}|g" "$file"
  rm "${file}.bak"
}


service_name=${1:-}
if [ -z "$service_name" ]; then
  echo "What service are you onboarding?"
  read -r service_name
fi
service_name=$(echo "$service_name" | tr '[:upper:]' '[:lower:]')
if ! [[ $service_name =~ ^[a-z0-9]+$ ]]; then
  echo "Invalid service name '${service_name}': use lowercase letters and digits only" >&2
  exit 1
fi
test ! -f "Makefile.${service_name}" || { echo "${service_name} already onboarded" >&2; exit 1; }

spec_url=${2:-}
if [ -z "$spec_url" ]; then
  echo "What is the URL for the OpenAPI spec for ${service_name}?"
  read -r spec_url
fi
spec_root_file=""

possible_spec_filename=${spec_url##*/}
if [[ $possible_spec_filename == *"."* ]]; then
  spec_root_file=$possible_spec_filename
  spec_url=${spec_url%/$spec_root_file}
fi
cp templates/Makefile.sdk Makefile.${service_name}
replace_in_file Makefile.${service_name} __PACKAGE_NAME__ "${service_name}"
replace_in_file Makefile.${service_name} __SPEC_BASE_URL__ "${spec_url}"
replace_in_file Makefile.${service_name} __SPEC_ROOT_FILE__ "${spec_root_file}"

mkdir -p spec/services/${service_name} && touch spec/services/${service_name}/.keep
mkdir -p templates/services/${service_name} && touch templates/services/${service_name}/.keep
mkdir -p patches/services/${service_name} && touch patches/services/${service_name}/.keep

if [ "${SKIP_GENERATE:-0}" != "1" ]; then
  make -f "Makefile.${service_name}" all
fi

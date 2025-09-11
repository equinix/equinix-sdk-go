#!/bin/bash

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


echo "What service are you onboarding?"
read service_name
service_name=$(echo "$service_name" | tr '[:upper:]' '[:lower:]')
test ! -f Makefile.${service_name} || (echo "${service_name} already onboarded" && exit 1)

echo "What is the URL for the OpenAPI spec for ${service_name}?"
read spec_url
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

cp templates/.github/workflows/sync.yaml .github/workflows/sync-${service_name}.yaml
replace_in_file .github/workflows/sync-${service_name}.yaml __PACKAGE_NAME__ "${service_name}"

mkdir -p spec/services/${service_name} && touch spec/services/${service_name}/.keep
mkdir -p templates/services/${service_name} && touch templates/services/${service_name}/.keep
mkdir -p patches/services/${service_name} && touch patches/services/${service_name}/.keep

make -f Makefile.${service_name} all

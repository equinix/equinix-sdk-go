#!/bin/bash

set -eu

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
sed -i '' "s/__PACKAGE_NAME__/${service_name}/g" Makefile.${service_name}
sed -i '' "s|__SPEC_BASE_URL__|${spec_url}|g" Makefile.${service_name}
sed -i '' "s|__SPEC_ROOT_FILE__|${spec_root_file}|g" Makefile.${service_name}

cp templates/.github/workflows/sync.yaml .github/workflows/sync-${service_name}.yaml
sed -i '' "s/__PACKAGE_NAME__/${service_name}/g" .github/workflows/sync-${service_name}.yaml

mkdir -p spec/services/${service_name} && touch spec/services/${service_name}/.keep
mkdir -p templates/services/${service_name} && touch templates/services/${service_name}/.keep
mkdir -p patches/services/${service_name} && touch patches/services/${service_name}/.keep

make -f Makefile.${service_name} all

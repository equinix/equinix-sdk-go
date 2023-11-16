#!/bin/sh

set -eu -o pipefail

function resolve_file () {
  local raw_file_path=$1
  if [ ! -f "$raw_file_path"  ]
  then
    mkdir -p $(dirname $raw_file_path)
    touch $raw_file_path
    file_path=$(realpath $raw_file_path)
    rm $raw_file_path
  else
    file_path=$(realpath $raw_file_path)
  fi
}

function resolve_relative_path () {
  local parent_path=$(realpath $1)
  resolve_file $2
  resolved_path=${file_path#$parent_path}
}

function download_file () {
  local spec_base_url=$1
  local output_dir=$2
  local path=$3
  local full_path=$output_dir$path
  if [ ! -f "$full_path"  ]
  then
    echo "Fetching $spec_base_url$path"

    local current_output_dir=$(dirname $full_path)
    mkdir -p $current_output_dir
    wget -O $full_path $spec_base_url$path

    ref_targets="$(yq e '.. | select(has("$ref")) | .$ref' $full_path)"

    for path in $ref_targets
    do
      resolve_relative_path $output_dir/ $current_output_dir/$path
      download_file $spec_base_url $output_dir $resolved_path
    done
  fi
}

SPEC_BASE_URL=$1
OUTPUT_BASE_DIR=$2
ROOT_FILE=$3

rm -rf $OUTPUT_BASE_DIR

download_file $SPEC_BASE_URL $OUTPUT_BASE_DIR /$ROOT_FILE

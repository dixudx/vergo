#!/usr/bin/env bash
# Copyright 2018 The Vergo Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

vergo::log::usage_from_stdin() {
  local messages=()
  while read -r line; do
    messages+=("$line")
  done

  vergo::log::usage "${messages[@]}"
}

# Print an usage message to stderr. The arguments are printed directly.
vergo::log::usage() {
  echo >&2
  local message
  for message; do
    echo "$message" >&2
  done
  echo >&2
}

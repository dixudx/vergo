#!/bin/bash
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

set -o errexit
set -o nounset
set -o pipefail

VERGO_ROOT=$(dirname "${BASH_SOURCE}")/..
source "${VERGO_ROOT}/hack/lib/init.sh"

cd "${VERGO_ROOT}"

if [[ -z "$(which gofmt)" ]]; then
    vergo::log::usage_from_stdin <<EOF
Can't find 'gofmt' in PATH, please fix and retry.
See https://golang.org/cmd/gofmt/ for more instructions.
EOF
    return 2
fi

find_go_files() {
  find . -not \( \
      \( \
        -wholename './output' \
        -o -wholename './_output' \
        -o -wholename './_gopath' \
        -o -wholename './release' \
        -o -wholename './target' \
        -o -wholename '*/third_party/*' \
        -o -wholename '*/vendor/*' \
      \) -prune \
    \) -name '*.go'
}

# gofmt exits with non-zero exit code if it finds a problem unrelated to
# formatting (e.g., a file does not parse correctly). Without "|| true" this
# would have led to no useful error message from gofmt, because the script would
# have failed before getting to the "echo" in the block below.
diff=$(find_go_files | xargs "$(which gofmt)" -d -s 2>&1) || true
if [[ -n "${diff}" ]]; then
  echo "${diff}"
  exit 1
fi

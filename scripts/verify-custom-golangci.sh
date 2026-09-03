#!/usr/bin/env sh
# Gostafa 2026.
# SPDX-License-Identifier: Apache-2.0.

set -eu

binary_path=${1:-./custom-golangci-lint}
binary_directory=$(cd "$(dirname "$binary_path")" && pwd)
binary="$binary_directory/$(basename "$binary_path")"
fixture=$(mktemp -d)

cleanup() {
  rm -rf "$fixture"
}

trap cleanup EXIT

mkdir -p "$fixture/module"

printf '%s\n' 'module example.com/vulner-fixture' 'go 1.24.0' > "$fixture/module/go.mod"
printf '%s\n' 'package fixture' 'func Run() {}' > "$fixture/module/fixture.go"
printf '%s\n' \
  'version: "2"' \
  'linters:' \
  '  default: none' \
  '  enable:' \
  '    - vulner' \
  '  settings:' \
  '    custom:' \
  '      vulner:' \
  '        type: module' \
  '        description: Report reachable Go vulnerabilities' \
  > "$fixture/module/.golangci.yml"

cd "$fixture/module"
"$binary" config verify
"$binary" run

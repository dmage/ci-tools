#!/usr/bin/env bash

set -euxo pipefail

echo "Running golangci-lint"
# CI has HOME set to '/' causing the linter to try and create a cache at /.cache for which
# it doesn't have permissions.
if [[ $HOME = '/' ]]; then
  export HOME=/tmp
fi

# We embedd this so it must exist for compilation to succeed, but it's not checked in
if [[ -n ${CI:-} ]]; then touch cmd/vault-secret-collection-manager/index.js; fi

golangci-lint run --build-tags e2e,e2e_framework,optional_operators

cd $(dirname $0)/..

# Make sure the failing linter doesn't make the command fail and hence the script pass because we don't enter
# the condition.
set +o pipefail
if go run ./vendor/github.com/polyfloyd/go-errorlint -errorf ./... 2>&1 \
  |grep 'non-wrapping format verb for fmt.Errorf'; then
  exit 1
fi

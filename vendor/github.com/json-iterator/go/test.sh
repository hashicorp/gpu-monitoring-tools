#!/usr/bin/env bash
# Copyright IBM Corp. 2018, 2020
# SPDX-License-Identifier: Apache-2.0


set -e
echo "" > coverage.txt

for d in $(go list ./... | grep -v vendor); do
    go test -coverprofile=profile.out -coverpkg=github.com/json-iterator/go $d
    if [ -f profile.out ]; then
        cat profile.out >> coverage.txt
        rm profile.out
    fi
done

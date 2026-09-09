#!/usr/bin/env bash
set -euo pipefail

VERSION="$(cat VERSION)"
COMMIT="$(git rev-parse --short HEAD)"
BUILD_DATE="$(date -u +%Y-%m-%dT%H:%M:%SZ)"

mkdir -p dist

go fmt ./...
go vet ./...
#go test ./...

OUTPUT="dist/tls-bridge-api"

CGO_ENABLED=0 go build -trimpath -ldflags="-s -w -X 'tls-bridge-api/version.Version=${VERSION}' -X 'tls-bridge-api/version.Commit=${COMMIT}' -X 'tls-bridge-api/version.BuildDate=${BUILD_DATE}'" -o "${OUTPUT}"


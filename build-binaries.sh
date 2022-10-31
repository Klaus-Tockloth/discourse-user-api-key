#!/bin/sh

# ------------------------------------
# Purpose:
# - Builds executables / binaries.
#
# Releases:
# - v1.0.0 - 2022-10-31: initial release
#
# Remarks:
# - go tool dist list
# ------------------------------------

# set -o xtrace
set -o verbose

# renew vendor content
go mod vendor

# compile 'aix'
env GOOS=aix GOARCH=ppc64 go build -o binaries/aix-ppc64/discourse-user-api-key

# compile 'darwin'
env GOOS=darwin GOARCH=amd64 go build -o binaries/darwin-amd64/discourse-user-api-key
env GOOS=darwin GOARCH=arm64 go build -o binaries/darwin-arm64/discourse-user-api-key

# compile 'dragonfly'
env GOOS=dragonfly GOARCH=amd64 go build -o binaries/dragonfly-amd64/discourse-user-api-key

# compile 'freebsd'
env GOOS=freebsd GOARCH=amd64 go build -o binaries/freebsd-amd64/discourse-user-api-key
env GOOS=freebsd GOARCH=arm64 go build -o binaries/freebsd-arm64/discourse-user-api-key

# compile 'illumos'
env GOOS=illumos GOARCH=amd64 go build -o binaries/illumos-amd64/discourse-user-api-key

# compile 'linux'
env GOOS=linux GOARCH=amd64 go build -o binaries/linux-amd64/discourse-user-api-key
env GOOS=linux GOARCH=arm64 go build -o binaries/linux-arm64/discourse-user-api-key
env GOOS=linux GOARCH=mips64 go build -o binaries/linux-mips64/discourse-user-api-key
env GOOS=linux GOARCH=mips64le go build -o binaries/linux-mips64le/discourse-user-api-key
env GOOS=linux GOARCH=ppc64 go build -o binaries/linux-ppc64/discourse-user-api-key
env GOOS=linux GOARCH=ppc64le go build -o binaries/linux-ppc64le/discourse-user-api-key
env GOOS=linux GOARCH=riscv64 go build -o binaries/linux-riscv64/discourse-user-api-key
env GOOS=linux GOARCH=s390x go build -o binaries/linux-s390x/discourse-user-api-key

# compile 'netbsd'
env GOOS=netbsd GOARCH=amd64 go build -o binaries/netbsd-amd64/discourse-user-api-key
env GOOS=netbsd GOARCH=arm64 go build -o binaries/netbsd-arm64/discourse-user-api-key

# compile 'openbsd'
env GOOS=openbsd GOARCH=amd64 go build -o binaries/openbsd-amd64/discourse-user-api-key
env GOOS=openbsd GOARCH=arm64 go build -o binaries/openbsd-arm64/discourse-user-api-key
env GOOS=openbsd GOARCH=mips64 go build -o binaries/openbsd-mips64/discourse-user-api-key

# compile 'solaris'
env GOOS=solaris GOARCH=amd64 go build -o binaries/solaris-amd64/discourse-user-api-key

# compile 'windows'
env GOOS=windows GOARCH=amd64 go build -o binaries/windows-amd64/discourse-user-api-key.exe
env GOOS=windows GOARCH=386 go build -o binaries/windows-386/discourse-user-api-key.exe
env GOOS=windows GOARCH=arm go build -o binaries/windows-arm/discourse-user-api-key.exe

#!/usr/bin/env bash

scriptDir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

go run ${scriptDir}/sdk/normalization/gen/*.go -out ${scriptDir}/sdk/normalization/event_functions.go || exit 1
gofmt -w ${scriptDir}/sdk/normalization/event_functions.go
go test ./...
go build -mod vendor -o /usr/local/bin/raccoon ${scriptDir}/main.go

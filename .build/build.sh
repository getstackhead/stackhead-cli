#!/bin/bash

go get github.com/markbates/pkger/cmd/pkger
pkger
go generate ./...
go build -o ./bin/stackhead-cli .

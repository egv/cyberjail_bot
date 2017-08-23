#!/bin/bash

docker run --rm -v "$PWD":/go/src/github.com/egv/cyberjail_bot -w /go/src/github.com/egv/cyberjail_bot -e GOOS=linux -e GOARCH=amd64 -e CGO_ENABLED=0 golang:1.8 go build -a -v

#!/usr/bin/env bash

export GOPATH=`pwd`
export GOBIN=`pwd`/bin

go clean
go install src/main/main.go

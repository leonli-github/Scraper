#!/usr/bin/env bash

export GOPATH=`pwd`

go clean
go run src/main/main.go $@


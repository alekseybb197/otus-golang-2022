#!/bin/bash

go fmt .
golangci-lint run .
go test -v -count=1 -race -timeout=1m .
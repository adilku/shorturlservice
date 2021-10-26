#!/bin/sh -l
go build -v ./cmd/shorturlservice && ./shorturlservice -config-path=./configs/dockershorturlservice.toml

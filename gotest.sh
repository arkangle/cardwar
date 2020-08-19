#!/bin/bash
set -e
gotestsum --format dots -- -coverprofile=coverage.out ./...
go tool cover -func=coverage.out

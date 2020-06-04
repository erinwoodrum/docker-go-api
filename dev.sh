#! /bin/bash

# Spins up database container, 
../build-docker.sh
# Set env var for fake db. 
source .env
# Runs with fresh `go run main.go`
go get github.com/pilu/fresh
fresh # watches for changes, and reruns go run main.go



#!/bin/bash

echo "Running gofmt..."
gofmt -l -w .

echo "Running go vet..."
go vet ./...

echo "Running golint..."
golint ./...

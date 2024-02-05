#!/bin/bash

echo "Running gofmt..."
# List & Write formatting differes and results to stdout
gofmt -l -w .

echo "Running go vet..."
go vet ./...

echo "Running golint..."
# Show as many warnings as possible (default threshold min_confidence=0.8)
golint -min_confidence=0.1 ./...

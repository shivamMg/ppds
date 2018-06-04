#!/bin/sh

GOFMT=$(gofmt -l .)
if [ -n "${GOFMT}" ]; then
    printf >&2 'gofmt failed for:\n%s\n' "${GOFMT}"
    exit 1
fi

GOLINT=$(golint ./... | grep -v -E '(Box.*|Arrow) should have comment')
if [ -n "${GOLINT}" ]; then
    printf >&2 'golint failed:\n%s\n' "${GOLINT}"
    exit 1
fi

GOVET=$(go tool vet . 2>&1)
if [ -n "${GOVET}" ]; then
    printf >&2 'go vet failed:\n%s\n' "${GOVET}"
    exit 1
fi

GOTEST=$(go test ./...)
if [ $? -ne 0 ]; then
    printf >&2 'go test failed:\n%s\n' "${GOTEST}"
    exit 1
fi

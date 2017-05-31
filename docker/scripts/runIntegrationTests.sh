#!/usr/bin/env bash

echo "==============: Running integration tests..."
go test -v $(glide novendor) --tags=integration
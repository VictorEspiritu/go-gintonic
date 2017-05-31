#!/usr/bin/env bash

echo "==============: Running unit tests..."
go test -v $(glide novendor)
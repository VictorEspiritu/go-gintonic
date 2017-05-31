#!/usr/bin/env bash

echo "==============: Removing cache..."
rm -rf .glide
rm -rf vendor
rm -rf dist
rm glide.lock

echo "==============: Creating folders && Changing permissions..."
mkdir vendor
mkdir .glide
chmod -R 777 vendor
chmod -R 777 .glide

echo "==============: Install glide dependencies..."
glide install

echo "==============: Update glide dependencies..."
glide update

echo "==============: Build..."
go build pkg/api/main/main.go
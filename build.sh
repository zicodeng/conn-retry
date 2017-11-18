#!/usr/bin/env bash
echo "building Linux executable..."
GOOS=linux go build
echo "building docker container image..."
docker build -t drstearns/mqconnect .
echo "cleaning up..."
go clean

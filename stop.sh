#!/usr/bin/env bash
docker rm -f gateway
docker rm -f mqsvr
docker network rm mqtest

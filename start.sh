#!/usr/bin/env bash
docker network create mqtest

# RabbitMQ container instance
docker run -d \
--network mqtest \
--hostname mqhost \
--name mqsvr \
rabbitmq

# API gateway container instance
docker run -d \
-p 80:80 \
--network mqtest \
--name gateway \
-e MQADDR=mqsvr \
drstearns/mqconnect

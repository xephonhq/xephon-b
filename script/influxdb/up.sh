#!/bin/sh

# https://hub.docker.com/_/influxdb/

docker run -p 8083:8083 -p 8086:8086 \
  influxdb

#!/bin/sh

# TODO: use docker compose
# This one use docke compose https://www.big-data-europe.eu/scalable-sparkhdfs-workbench-using-docker/
# The example used in hadoop's doc https://github.com/sequenceiq/hadoop-docker
# https://github.com/krejcmat/hadoop-hbase-docker
# https://github.com/dajobe/hbase-docker
# My old scripts for setting up all the Hxx https://github.com/at15/dev-node

# https://hub.docker.com/r/petergrace/opentsdb-docker/
docker run -p 4242:4242 \
  petergrace/opentsdb-docker:latest
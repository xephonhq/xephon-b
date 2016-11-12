#!/bin/sh

# https://hub.docker.com/_/elasticsearch/


sudo sysctl -w vm.max_map_count=262144

# use 1/2 of available memory
avail_mem=`free -m | grep Mem | awk '{print $7}'`
jvm_mem=$((avail_mem / 2))

docker run -p 9200:9200 -p 9300:9300 \
  -e ES_JAVA_OPTS="-Xms${jvm_mem}m -Xmx${jvm_mem}m" \
  elasticsearch:5.0

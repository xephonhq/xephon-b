#!/bin/sh

# https://hub.docker.com/_/elasticsearch/

c=$(cat /proc/sys/vm/max_map_count)
echo "vm.max_map_count is ${c}"

if [ "$c" -lt "262144" ]; then
    echo "vm.max_map_count must be bigger than 262144 https://www.elastic.co/guide/en/elasticsearch/reference/5.0/_maximum_map_count_check.html"
    echo "sysctl -w vm.max_map_count=262144"
    sudo sysctl -w vm.max_map_count=262144
fi

# use 1/2 of available memory https://www.elastic.co/guide/en/elasticsearch/guide/current/heap-sizing.html
avail_mem=`free -m | grep Mem | awk '{print $7}'`
jvm_mem=$((avail_mem / 2))
echo "available memory is ${avail_mem}, give half to jvm, which is ${jvm_mem}"

docker run -p 9200:9200 -p 9300:9300 \
  -e ES_JAVA_OPTS="-Xms${jvm_mem}m -Xmx${jvm_mem}m" \
  elasticsearch:5.0

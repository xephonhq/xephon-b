FROM openjdk:8-stretch
# NOTE: we no longer use alphine because it won't run program with cgo and our waitfor it is copied from local env directly
# where cgo is enabled, this makes the image much larger ...

MAINTAINER at15 at15@dongyue.io

RUN mkdir /opt; \
  cd /opt; \
  curl -L https://github.com/kairosdb/kairosdb/releases/download/v1.2.0/kairosdb-1.2.0-1.tar.gz | \
  tar zxvfp -

# Use Cassandra
COPY kairosdb.properties /opt/kairosdb/conf/kairosdb.properties

EXPOSE 4242 8080 2003 2004

ENTRYPOINT []
CMD [ "/opt/kairosdb/bin/kairosdb.sh", "run"]

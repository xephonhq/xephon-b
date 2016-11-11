# KairosDB

- http://kairosdb.github.io/

## Usage

Build and run a single KairosDB node with H2

TODO: need to manually change to H2 in kairosdb.properties

- `docker build -t xephonhq/kairosdb ./node`
- `docker run -p 8080:8080 --name xephonhq-kairosdb xephonhq/kairosdb`

## Requirement

- [official docker support for Kairosdb?](https://github.com/kairosdb/kairosdb/issues/288)
- JDK7/8
- Cassandra 2.2

## Docker images

- https://github.com/cit-lab/kairosdb/tree/feature/alpine

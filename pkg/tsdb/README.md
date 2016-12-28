# TSDB

This directory hold

- unified data type, query interface
- the client code for various TSDBs
- the server API implementation for various TSDBs

## Client

- Client is actually a bunch of client connections with TSDB which share same transport, see initial discussion [here](https://github.com/xephonhq/xephon-b/pull/14)
  - [ ] TODO: may have an abstraction over http client 
- You can control trace (collect client metrics), concurrency, retry policy
- Payload is not thread safe, if you want to fine control over payload, use lock on it.
  - [ ] TODO: add lock to payload struct
- Add points to client is thread safe, and how the client handle it is based on config,
  i.e. it may batch it until there are enough points


## Payload 

- It acts like what is called `builder` in some TSDB Java clients, you can add points to it and they get bytes.
- It is NOT thread safe. 
- It has two modes 

### Simple 

- Add points without any grouping, it will turn every point into bytes directly
- Just call `Bytes()`, once you call `Bytes()` you can NO longer add points to the payload, and there is no good reason for doing that.

### Group by series 

- When adding a point with series, its meta will be sorted, if the same meta already exists, the point will be add to it. 
Otherwise `SeriesWithIntPoints` is added.
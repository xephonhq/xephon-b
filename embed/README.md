# Embed

Embed test tsdb storage engine directly without going through wire protocol, so the test is more accurate.
i.e. you are benchmarking the compression, cache etc. instead of http/tcp server.

## Known issue

- need to clone influxdb and prometheus to GOPATH because we use replace in go mod to avoid all the errors ...

## References

- https://github.com/influxdata/influxdb/tree/master/query/promql/internal/promqltests
- https://github.com/prometheus/prombench
- https://github.com/prometheus/prometheus/blob/master/tsdb/cmd/tsdb/main.go
- https://github.com/smallnest/kvbench Server for benchmarking pure Go key/value databases
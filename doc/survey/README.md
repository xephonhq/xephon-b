# Survey

## Http Benchmark

- [ ] how does client know it is using http 1.1 or http 2.0
- [x] https://github.com/rakyll/hey
  - it had the max connection issue https://github.com/rakyll/hey/issues/31, didn't reuse transport?
  - [Feature : Distributed load sending](https://github.com/rakyll/hey/issues/91) use mqtt
- [x] https://github.com/adjust/go-wrk
  - support leader and follower using simple http API
- [ ] wrk
- [ ] wrk with HDR Histogram
- [ ] istio fortio
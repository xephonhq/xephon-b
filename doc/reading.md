# Reading

Mainly about the following three areas

- Time series database design
- Benchmark tool design
- Distributed database design

## Time series databases design

- Cassandra http://www.datastax.com/dev/blog/advanced-time-series-with-cassandra
- Gorilla http://www.vldb.org/pvldb/vol8/p1816-teller.pdf
- A survey blog http://jmoiron.net/blog/thoughts-on-timeseries-databases

## Benchmark tool design

### Time series database benchmark tools

- From InfluxDB https://github.com/influxdata/influxdb-comparisons, [data generation](https://github.com/influxdata/influxdb-comparisons/tree/master/bulk_data_gen)

### YCSB

- https://github.com/brianfrankcooper/YCSB
- [Benchmarking cloud serving systems with YCSB](http://dl.acm.org/citation.cfm?id=1807152) published on [SoCC 2010](http://research.microsoft.com/en-us/um/redmond/events/socc2010/index.htm)

#### Publications that referred YCSB

- [A Model-based Approach to Database Stress Testing](http://orbilu.uni.lu/bitstream/10993/28616/1/preprint_dexa2016.pdf) in which they found a bug in
[VoltDB](https://www.voltdb.com/)

# Berkeley Tree Database (BTrDB)

## Paper

- [BTrDB: Optimizing Storage System Design for Timeseries Processing](https://www.usenix.org/system/files/conference/fast16/fast16-papers-andersen.pdf)
- [Review from at15](https://github.com/at15/papers-i-read/blob/master/databases/btrdb.md)

## Code

- https://github.com/SoftwareDefinedBuildings/btrdb

## Benchmark

- https://github.com/SoftwareDefinedBuildings/btrdb-test/blob/master/loadgen/loadgen.go

## Survey

- KairosDB is faster than OpenTSDB
- aggregate
  - on the fly: OpenTSDB, Druid
  - pre computed aggregates: InfluxDB, RespawnDB
  - [ ] I think BTrDB is kind of pre computed since it uses a tree structure
- 'many good solutions exist for querying metadata to obtain a collection of streams.' so only focus on stream

## Special

- high throughout that all existing tsdb can not handle from telemetry (microsynchophasors, or uPMUs)
- high precision, nanoseconds
- support out order of arrival, duplicate data
  - 'raw data streams feed into a graph of distillation processes in order to clean and filter the raw data' but the following sentence is quite confusing 'In
  presence of out of order arrival and loss, without support from the storage engine, it can be complex and costly to determine which input ranges have changed
  and which output extents need to be computed, or recomputed, to maintain consistency throughout the distillation pipeline'. (it seems its for the ComputeDiff method)
- Time partitioning copy on write version-annotated k-ary tree (COW in the paper)
  - old version can be visited when new version is updated
  - 'The tree need only be walked to the depth of the desired difference resolution'. data is pre-aggregated in parent node (min, max, mean, count), so only parent nodes are visited when low resolution is needed. even the resolution is not 2 powered, only the two end need to go down to deeper level.
  - the tree can cover time in nanoseconds from 1933 to 2079
- [ ] **eventual consistency** in a graph of independent analytics despite out of order or duplicate data (still didn't quite get it, and it is not well defined in the paper)

## Can be used

- [SEDA: An Architecture for Well-Conditioned, Scalable Internet Services ](http://www.sosp.org/2001/papers/welsh.pdf) We write our application following this bla bla

## Reference

Some reference that I found pretty useful

- [Solving big data challenges for enterprise application performance management](http://vldb.org/pvldb/vol5/p1724_tilmannrabl_vldb2012.pdf) they use
YCSB to simulate APM (Application performance monitor data) on Voldemort, Redis, HBase, Cassandra, MySQL and VoltDB **We can argue that their measurement is
only cover a small part of time series database use**
- [Respawn: A Distributed Multi-Resolution Time-Series Datastore](https://users.ece.cmu.edu/~agr/resources/publications/respawn-rtss-13.pdf) pre computed aggregates

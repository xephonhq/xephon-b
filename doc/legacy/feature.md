# Features 

The Xephon-B benchmark suite contains three parts

- workload generator and executor
- probe on database machine which collect and store metrics of hardware and the database itself.
- benchmark result store and Web UI from which people can query by their interest and compare result from different sources (ie: database vendors, researchers)

## Workload generator and executor

Our key features are

- Flexibility and Extensibility, support multiple time series databases and easy to add new one.
- (Write) Time series specific data generator
- (Read) Single tenant and Multi tenant

We simulate two real world scenarios

- Monitoring large scale distributed systems.
- Smart health device with end users querying their health data.

## Probe on database machine

Our key features are

- Collect machine and database metrics to combine with metrics from workload executor.
- Store the data in local time series database.
- Simulate some hardware and network failures.
- Use [Consul](https://www.consul.io/) (based on Raft) for service discovery and configuration.

## Benchmark result store and Web UI

Our key features are

- A central store for multiple sources benchmark results, researchers, companies can publish their raw benchmark data and share their configuration.
- Compare against different versions of a same database, vendors can show how they improve the performance.
- Powerful query and graphing.
- Link to publications and blog posts to strengthen your points.

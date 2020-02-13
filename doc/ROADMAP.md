# Roadmap

## Up coming

### v0.4.0

Test compress ratio and compress speed, value generator is likely going to be migrated to libtsdb-go

- constant value (should show those using RLE is better than xor alike)
- fixed delta
- fixed double delta
- random delta
- distribution?
- real monitoring data (there is a paper that says several data source ...)

### v0.5.0

Test some popular storage engines

- influxdb
- prometheus
- m3db (if we can find a way to test its single node)
- victoria metrics
- xephon-k or libtsdb-go, if we can write it out ...

## Finished


# KairosDB

## Design

- [Cassandra Schema](https://kairosdb.github.io/docs/build/html/CassandraSchema.html)

The precision is millisecond. 3 x 7 x 24 x 60 x 60 x 1000 = 1,814,400,000
> The length of the row is set to exactly three weeks of data or 1,814,400,000 columns.

The design is similar to the blog post, use a meta table to keep where certain row starts

## Code

- https://github.com/kairosdb/kairosdb

## Setup

- `docker run --name kairosdb-cassandra -e CASSANDRA_START_RPC=true -p 9160:9160 -d cassandra:2.2` and turn on thrift
- `kairosdb.service.datastore=org.kairosdb.datastore.cassandra.CassandraModule`
- [ ] still can't talk to the thrift protocol .... the config file in container is right
- http://cassandra.apache.org/doc/latest/getting_started/installing.html#installation-from-binary-tarball-files (use local cassandra works)

## Schema

````java
private void createSchema(int replicationFactor)
{
    List<ColumnFamilyDefinition> cfDef = new ArrayList<ColumnFamilyDefinition>();

    cfDef.add(HFactory.createColumnFamilyDefinition(
            m_keyspaceName, CF_DATA_POINTS, ComparatorType.BYTESTYPE));

    cfDef.add(HFactory.createColumnFamilyDefinition(
            m_keyspaceName, CF_ROW_KEY_INDEX, ComparatorType.BYTESTYPE));

    cfDef.add(HFactory.createColumnFamilyDefinition(
            m_keyspaceName, CF_STRING_INDEX, ComparatorType.UTF8TYPE));

    KeyspaceDefinition newKeyspace = HFactory.createKeyspaceDefinition(
            m_keyspaceName, ThriftKsDef.DEF_STRATEGY_CLASS,
            replicationFactor, cfDef);

    m_cluster.addKeyspace(newKeyspace, true);
}
````
- `select * from system.schema_keyspaces;` or `describe keyspaces`
- `use kairosdb`
- `describe tables`
- `describe table string_index`

data_points
````
CREATE TABLE kairosdb.data_points (
    key blob,
    column1 blob,
    value blob,
    PRIMARY KEY (key, column1)
) WITH COMPACT STORAGE
    AND CLUSTERING ORDER BY (column1 ASC)
    AND bloom_filter_fp_chance = 0.01
    AND caching = '{"keys":"ALL", "rows_per_partition":"NONE"}'
    AND comment = ''
    AND compaction = {'class': 'org.apache.cassandra.db.compaction.SizeTieredCompactionStrategy'}
    AND compression = {'sstable_compression': 'org.apache.cassandra.io.compress.LZ4Compressor'}
    AND dclocal_read_repair_chance = 0.1
    AND default_time_to_live = 0
    AND gc_grace_seconds = 864000
    AND max_index_interval = 2048
    AND memtable_flush_period_in_ms = 0
    AND min_index_interval = 128
    AND read_repair_chance = 1.0
    AND speculative_retry = 'NONE';

````
row_key_index
````
CREATE TABLE kairosdb.row_key_index (
    key blob,
    column1 blob,
    value blob,
    PRIMARY KEY (key, column1)
) WITH COMPACT STORAGE
    AND CLUSTERING ORDER BY (column1 ASC)
    AND bloom_filter_fp_chance = 0.01
    AND caching = '{"keys":"ALL", "rows_per_partition":"NONE"}'
    AND comment = ''
    AND compaction = {'class': 'org.apache.cassandra.db.compaction.SizeTieredCompactionStrategy'}
    AND compression = {'sstable_compression': 'org.apache.cassandra.io.compress.LZ4Compressor'}
    AND dclocal_read_repair_chance = 0.1
    AND default_time_to_live = 0
    AND gc_grace_seconds = 864000
    AND max_index_interval = 2048
    AND memtable_flush_period_in_ms = 0
    AND min_index_interval = 128
    AND read_repair_chance = 1.0
    AND speculative_retry = 'NONE';
````
string_index
````
CREATE TABLE kairosdb.string_index (
    key blob,
    column1 text,
    value blob,
    PRIMARY KEY (key, column1)
) WITH COMPACT STORAGE
    AND CLUSTERING ORDER BY (column1 ASC)
    AND bloom_filter_fp_chance = 0.01
    AND caching = '{"keys":"ALL", "rows_per_partition":"NONE"}'
    AND comment = ''
    AND compaction = {'class': 'org.apache.cassandra.db.compaction.SizeTieredCompactionStrategy'}
    AND compression = {'sstable_compression': 'org.apache.cassandra.io.compress.LZ4Compressor'}
    AND dclocal_read_repair_chance = 0.1
    AND default_time_to_live = 0
    AND gc_grace_seconds = 864000
    AND max_index_interval = 2048
    AND memtable_flush_period_in_ms = 0
    AND min_index_interval = 128
    AND read_repair_chance = 1.0
    AND speculative_retry = 'NONE';
````

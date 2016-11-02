# KairosDB

## Design

- [Cassandra Schema](https://kairosdb.github.io/docs/build/html/CassandraSchema.html)

The precision is millisecond. 3 x 7 x 24 x 60 x 60 x 1000 = 1,814,400,000
> The length of the row is set to exactly three weeks of data or 1,814,400,000 columns.

The design is similar to the blog post, use a meta table to keep where certain row starts

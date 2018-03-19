package metrics

import "github.com/libtsdb/libtsdb-go/libtsdb"

type Response interface {
	libtsdb.Trace
}

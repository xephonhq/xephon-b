package metrics

import "github.com/libtsdb/libtsdb-go/protocol"

type Response interface {
	protocol.Trace
}

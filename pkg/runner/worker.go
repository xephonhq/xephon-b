package runner

import (
	"github.com/libtsdb/libtsdb-go/libtsdb"
)

type Worker struct {
	c libtsdb.WriteClient
}

package runner

import (
	"context"

	dlog "github.com/dyweb/gommon/log"

	"github.com/libtsdb/libtsdb-go/libtsdb"
	"github.com/xephonhq/xephon-b/pkg/config"
)

type Worker struct {
	cfg config.WorkloadConfig
	c   libtsdb.WriteClient
	// TODO: each worker should have a generator
	// TODO: reporter
	log *dlog.Logger
}

func NewWorker(c libtsdb.WriteClient) (*Worker, error) {
	w := &Worker{
		c: c,
	}
	dlog.NewStructLogger(log, w)
	return w, nil
}

func (w *Worker) Run(ctx context.Context) error {
	return nil
}

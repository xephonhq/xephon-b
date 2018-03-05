package runner

import (
	"context"

	"github.com/dyweb/gommon/errors"
	dlog "github.com/dyweb/gommon/log"

	"github.com/libtsdb/libtsdb-go/libtsdb"
	"github.com/libtsdb/libtsdb-go/libtsdb/client/graphitew"
	"github.com/libtsdb/libtsdb-go/libtsdb/client/influxdbw"
	"github.com/libtsdb/libtsdb-go/libtsdb/client/kairosdbw"
	"github.com/xephonhq/xephon-b/pkg/config"
)

type Worker struct {
	wcfg config.WorkloadConfig
	dcfg config.DatabaseConfig
	c    libtsdb.WriteClient
	// TODO: each worker should have a generator
	// TODO: reporter
	log *dlog.Logger
}

func NewWorker(wcfg config.WorkloadConfig, dcfg config.DatabaseConfig) (*Worker, error) {
	w := &Worker{
		wcfg: wcfg,
		dcfg: dcfg,
	}
	dlog.NewStructLogger(log, w)
	return w, nil
}

func (w *Worker) Run(ctx context.Context) error {
	w.log.Infof("TODO: worker should do something")
	return nil
}

func createClient(cfg config.DatabaseConfig) (libtsdb.WriteClient, error) {
	switch cfg.Type {
	case "influxdb":
		return influxdbw.New(*cfg.Influxdb)
	case "kairosdb":
		return kairosdbw.New(*cfg.Kairosdb)
	case "graphite":
		return graphitew.New(*cfg.Graphite)
	default:
		return nil, errors.Errorf("unknown databse %s", cfg.Type)
	}
}

package runner

import (
	"context"
	"time"

	"github.com/dyweb/gommon/errors"
	dlog "github.com/dyweb/gommon/log"

	"github.com/libtsdb/libtsdb-go/libtsdb"
	"github.com/libtsdb/libtsdb-go/libtsdb/client/graphitew"
	"github.com/libtsdb/libtsdb-go/libtsdb/client/influxdbw"
	"github.com/libtsdb/libtsdb-go/libtsdb/client/kairosdbw"

	"github.com/xephonhq/xephon-b/pkg/config"
	"github.com/xephonhq/xephon-b/pkg/generator"
)

type Worker struct {
	// TODO: each worker should have a generator
	// TODO: reporter
	id   int
	wcfg config.WorkloadConfig
	dcfg config.DatabaseConfig
	c    libtsdb.WriteClient
	t    generator.TimeGenerator
	v    generator.ValueGenerator

	log *dlog.Logger
}

func NewWorker(id int, wcfg config.WorkloadConfig, dcfg config.DatabaseConfig) (*Worker, error) {
	c, err := createClient(dcfg)
	if err != nil {
		return nil, err
	}
	t, err := createTimeGenerator(wcfg.Time, c.Meta().TimePrecision)
	if err != nil {
		return nil, err
	}
	v, err := createValueGenerator(wcfg.Value)
	if err != nil {
		return nil, err
	}
	w := &Worker{
		id:   id,
		wcfg: wcfg,
		dcfg: dcfg,
		c:    c,
		t:    t,
		v:    v,
	}
	dlog.NewStructLogger(log, w)
	return w, nil
}

func (w *Worker) Run(ctx context.Context) error {
	w.log.Infof("TODO: worker should do something")
	for {
		select {
		case <-ctx.Done():
			log.Infof("worker %d exit due to context", w.id)
			return nil
		default:
			time.Sleep(time.Second)
		}
	}
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
		return nil, errors.Errorf("unknown database %s", cfg.Type)
	}
}

func createTimeGenerator(cfg config.TimeConfig, precision time.Duration) (generator.TimeGenerator, error) {
	return generator.NewFixIntervalTime(cfg.Interval, precision)
}

func createValueGenerator(cfg config.ValueConfig) (generator.ValueGenerator, error) {
	switch cfg.Generator {
	case "constant":
		return generator.NewConstant(cfg.Constant.Int, cfg.Constant.Double), nil
	default:
		return nil, errors.Errorf("unknown generator %s", cfg.Generator)
	}
}

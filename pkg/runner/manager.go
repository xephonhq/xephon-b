package runner

import (
	"context"
	"github.com/dyweb/gommon/errors"
	dlog "github.com/dyweb/gommon/log"
	"sync"

	"github.com/libtsdb/libtsdb-go/libtsdb"
	"github.com/libtsdb/libtsdb-go/libtsdb/client/graphitew"
	"github.com/libtsdb/libtsdb-go/libtsdb/client/influxdbw"
	"github.com/libtsdb/libtsdb-go/libtsdb/client/kairosdbw"

	"github.com/xephonhq/xephon-b/pkg/config"
)

type Manager struct {
	cfg config.XephonBConfig
	log *dlog.Logger
}

func NewManager(cfg config.XephonBConfig) (*Manager, error) {
	m := &Manager{
		cfg: cfg,
	}
	dlog.NewStructLogger(log, m)
	return m, nil
}

func (m *Manager) Run(ctx context.Context) error {
	cfg := m.cfg

	// set stop condition
	ctx, cancel := context.WithCancel(ctx)
	switch cfg.Limit {
	case "time":
		ctx, cancel = context.WithTimeout(ctx, cfg.Duration)
	case "points":
		// noop
	default:
		return errors.Errorf("unknown limit %s", cfg.Limit)
	}

	// read database config,
	var dbcfg *config.DatabaseConfig
	for i := range cfg.Databases {
		c := cfg.Databases[i]
		if c.Name == cfg.Database {
			m.log.Infof("target database is %s type %s", c.Name, c.Type)
			dbcfg = &c
			break
		}
	}
	if dbcfg == nil {
		return errors.Errorf("databse %s does not have config, check name in databases section", cfg.Database)
	}

	var wg sync.WaitGroup
	// TODO: first start reporter

	// worker
	if cfg.Worker.Num <= 0 {
		return errors.Errorf("invalid worker number %d", cfg.Worker.Num)
	}
	for i := 0; i < cfg.Worker.Num; i++ {
		c, err := createClient(*dbcfg)
		if err != nil {
			m.log.Errorf("can't create tsdb client use config %s", err.Error())
			cancel()
			break
		}
		wg.Add(1)
		go func(c libtsdb.WriteClient) {
			m.log.Infof("TODO: make request")
			wg.Done()
		}(c)
	}
	wg.Wait()
	cancel()
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

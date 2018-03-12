package runner

import (
	"context"
	"fmt"
	"sync"

	"github.com/dyweb/gommon/errors"
	dlog "github.com/dyweb/gommon/log"
	"github.com/xephonhq/xephon-b/pkg/config"
	"github.com/xephonhq/xephon-b/pkg/metrics"
	"github.com/xephonhq/xephon-b/pkg/reporter"
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
		// FIXME: limit by points is not implemented
	default:
		return errors.Errorf("unknown limit %s", cfg.Limit)
	}

	// read config
	var (
		dbcfg config.DatabaseConfig
		wlcfg config.WorkloadConfig
		rpcfg config.ReporterConfig
		err   error
	)
	if dbcfg, err = m.selectDatabase(); err != nil {
		return err
	}
	if wlcfg, err = m.selectWorkload(); err != nil {
		return err
	}
	if rpcfg, err = m.selectReporter(); err != nil {
		return err
	}

	var wg sync.WaitGroup
	resChan := make(chan metrics.Response, cfg.Worker.Num)

	// reporter
	rep, err := createReporter(rpcfg)
	if err != nil {
		return err
	}
	repCtx, repCancel := context.WithCancel(ctx)
	go func() {
		rep.Run(repCtx, resChan)
	}()
	// worker
	if cfg.Worker.Num <= 0 {
		return errors.Errorf("invalid worker number %d", cfg.Worker.Num)
	}
	// create workers, exit if any of them has error
	workers := make([]*Worker, cfg.Worker.Num)
	for i := 0; i < cfg.Worker.Num; i++ {
		if wk, err := NewWorker(i, wlcfg, dbcfg, resChan); err != nil {
			return errors.Wrap(err, "can't create worker")
		} else {
			workers[i] = wk
		}
	}
	// run workers
	for i := 0; i < cfg.Worker.Num; i++ {
		wg.Add(1)
		go func(wk *Worker) {
			// TODO: cancel when error
			wk.Run(ctx)
			wg.Done()
		}(workers[i])
	}
	wg.Wait()
	cancel()
	repCancel()
	if err := rep.Finalize(); err != nil {
		return errors.Wrap(err, "can't finalize reporter")
	}
	if err := rep.Flush(); err != nil {
		return errors.Wrap(err, "can't flush reporter")
	}
	fmt.Println(rep.TextReport())
	// TODO: write text and json report to somewhere ...
	return nil
}

func (m *Manager) selectDatabase() (config.DatabaseConfig, error) {
	for _, c := range m.cfg.Databases {
		if c.Name == m.cfg.Database {
			m.log.Infof("target database is %s type %s", c.Name, c.Type)
			return c, nil
		}
	}
	return config.DatabaseConfig{},
		errors.Errorf("database %s does not have config, check name in databases section", m.cfg.Database)
}

func (m *Manager) selectWorkload() (config.WorkloadConfig, error) {
	for _, c := range m.cfg.Workloads {
		if c.Name == m.cfg.Workload {
			m.log.Infof("workload is %s series %d value generator is %v", c.Name, c.Series.Num, c.Value.Generator)
			return c, nil
		}
	}
	return config.WorkloadConfig{},
		errors.Errorf("workload %s does not have config, check name in workloads section", m.cfg.Workload)
}

func (m *Manager) selectReporter() (config.ReporterConfig, error) {
	for _, c := range m.cfg.Reporters {
		if c.Name == m.cfg.Reporter {
			m.log.Infof("reporter %s is type %s", c.Name, c.Type)
			return c, nil
		}
	}
	return config.ReporterConfig{},
		errors.Errorf("reporter %s does not have config, check name in reporters section", m.cfg.Reporter)
}

func createReporter(cfg config.ReporterConfig) (reporter.Sink, error) {
	switch cfg.Type {
	// TODO: define string as constant in config package ReporterTypeCounter etc.
	case "counter":
		if cfg.Counter == nil {
			return nil, errors.Errorf("counter is selected but no config")
		}
		return reporter.NewCounter(*cfg.Counter), nil
	}
	return nil, errors.Errorf("unknown reporter type %s", cfg.Type)
}

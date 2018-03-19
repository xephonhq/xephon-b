package runner

import (
	"context"
	"time"

	"github.com/dyweb/gommon/errors"
	dlog "github.com/dyweb/gommon/log"

	"github.com/libtsdb/libtsdb-go/libtsdb"

	pb "github.com/libtsdb/libtsdb-go/libtsdb/libtsdbpb"

	"github.com/xephonhq/xephon-b/pkg/config"
	"github.com/xephonhq/xephon-b/pkg/generator"
	"github.com/xephonhq/xephon-b/pkg/metrics"
	"github.com/xephonhq/xephon-b/pkg/util/tsdbutil"
)

type Worker struct {
	id   int
	wcfg config.WorkloadConfig
	dcfg config.DatabaseConfig
	c    libtsdb.TracedWriteClient

	// generator
	sGen generator.SeriesGenerator
	tGen generator.TimeGenerator
	vGen generator.ValueGenerator

	// reporter
	resChan chan<- metrics.Response

	log *dlog.Logger
}

func NewWorker(id int,
	wcfg config.WorkloadConfig, dcfg config.DatabaseConfig,
	resChan chan<- metrics.Response) (*Worker, error) {
	// check workload config
	if wcfg.Batch.Series <= 0 || wcfg.Batch.Points <= 0 {
		return nil, errors.Errorf("invalid batch series %d or points %d", wcfg.Batch.Series, wcfg.Batch.Points)
	}
	c, err := tsdbutil.CreateClient(dcfg)
	if err != nil {
		return nil, err
	}
	s, err := createSeriesGenerator(wcfg.Series)
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
		id:      id,
		wcfg:    wcfg,
		dcfg:    dcfg,
		c:       c,
		sGen:    s,
		tGen:    t,
		vGen:    v,
		resChan: resChan,
	}
	dlog.NewStructLogger(log, w)
	return w, nil
}

func (w *Worker) Run(ctx context.Context) error {
	w.log.Infof("worker %d started", w.id)
	for {
		select {
		case <-ctx.Done():
			log.Infof("worker %d exit due to context", w.id)
			return nil
		default:
			w.genBatch()
			err := w.c.Flush()
			if err != nil {
				log.Warnf("failed to flush %s", err.Error())
			}
			w.resChan <- w.c.Trace()
		}
	}
	return nil
}

func (w *Worker) genBatch() {
	t := w.tGen.NextTime()
	for i := 0; i < w.wcfg.Batch.Series; i++ {
		sMeta := w.sGen.NextSeries()
		for j := 0; j < w.wcfg.Batch.Points; j++ {
			// FIXME: we hardcoded to use float, should allow mix them or at least pick one ...
			v := w.vGen.NextDouble()
			p := pb.PointDoubleTagged{
				Name: sMeta.Name,
				Tags: sMeta.Tags,
				Point: pb.PointDouble{
					T: t,
					V: v,
				},
			}
			// TODO: we should put many points in a series for tsdb that supports this, i.e. KairosDB, OpenTSDB, it is not supported by libtsdb yet
			w.c.WriteDoublePoint(&p)
		}
	}
}

func createSeriesGenerator(cfg config.SeriesConfig) (generator.SeriesGenerator, error) {
	// TODO: allow more than one g
	return generator.NewGenericSeries(cfg)
}

func createTimeGenerator(cfg config.TimeConfig, precision time.Duration) (generator.TimeGenerator, error) {
	return generator.NewFixIntervalTime(cfg.Interval, precision)
}

func createValueGenerator(cfg config.ValueConfig) (generator.ValueGenerator, error) {
	switch cfg.Generator {
	case "constant":
		return generator.NewConstant(cfg.Constant.Int, cfg.Constant.Double), nil
	case "random":
		// TODO: apply config, min, max
		return generator.NewRandom(), nil
	default:
		return nil, errors.Errorf("unknown generator %s", cfg.Generator)
	}
}

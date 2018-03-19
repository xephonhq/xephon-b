package reporter

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/dyweb/gommon/errors"
	dlog "github.com/dyweb/gommon/log"
	"github.com/libtsdb/libtsdb-go/libtsdb"
	pb "github.com/libtsdb/libtsdb-go/libtsdb/libtsdbpb"
	"github.com/xephonhq/xephon-b/pkg/config"
	"github.com/xephonhq/xephon-b/pkg/metrics"
	"github.com/xephonhq/xephon-b/pkg/util/tsdbutil"
)

var _ Sink = (*TSDB)(nil)

// TSDB writes results as time series
type TSDB struct {
	cfg          config.TSDBReporterConfig
	globalConfig config.XephonBConfig
	c            libtsdb.WriteClient
	precision    time.Duration

	bufferedPoints int
	startTime      time.Time
	endTime        time.Time

	log *dlog.Logger
}

func NewTSDB(cfg config.TSDBReporterConfig, gcfg config.XephonBConfig) (*TSDB, error) {
	c, err := tsdbutil.CreateClient(cfg.Database)
	if err != nil {
		return nil, errors.Wrap(err, "need tsdb client")
	}
	tsdb := &TSDB{
		cfg:          cfg,
		globalConfig: gcfg,
		c:            c,
		precision:    c.Meta().TimePrecision,
	}
	dlog.NewStructLogger(log, tsdb)
	return tsdb, nil
}

func (d *TSDB) Run(ctx context.Context, resCh <-chan metrics.Response) {
	d.startTime = time.Now()
	d.log.Infof("tsdb reporter start %s", d.startTime)
	for {
		select {
		case <-ctx.Done():
			d.log.Info("tsdb reporter stopped by context")
			goto END
		case res, ok := <-resCh:
			if !ok {
				d.log.Info("tsdb reporter stopped by closed channel")
				goto END
			}
			d.Record(res)
		}
	}
END:
	d.endTime = time.Now()
	d.log.Infof("tsdb reporter stop %s duration %s", d.endTime, d.endTime.Sub(d.startTime))
}

func (d *TSDB) Record(res metrics.Response) {
	// NOTE: only keep track of latency
	// FIXME: support sampling
	// FIXME: allow config flush size
	t := time.Now().UnixNano()
	if d.precision == time.Second {
		t = t / (1000 * 1000 * 1000)
	} else if d.precision == time.Millisecond {
		t = t / (1000 * 100)
	} else if d.precision == time.Microsecond {
		t = t / 1000
	}
	// TODO: distinguish worker ...
	p := pb.PointIntTagged{
		Name: "latency",
		Tags: []pb.Tag{
			{K: "db", V: d.globalConfig.Database},
			// TODO: actually response should pass worker id so they are stored in different series
			{K: "totalWorker", V: fmt.Sprintf("%d", d.globalConfig.Worker.Num)},
			{K: "workload", V: d.globalConfig.Workload},
		},
		Point: pb.PointInt{
			T: t,
			V: res.GetEndTime() - res.GetStartTime(),
		},
	}
	d.c.WriteIntPoint(&p)
	d.bufferedPoints++
	if d.bufferedPoints > 100 {
		d.bufferedPoints = 0
		if err := d.c.Flush(); err != nil {
			d.log.Warnf("tsdb reporter flush failed %v", err)
		}
	}
}

func (d *TSDB) Finalize() error {
	d.log.Info("finalize tsdb reporter, nothing to do")
	return nil
}

func (d *TSDB) Flush() error {
	d.log.Info("flush tsdb reporter")
	return d.c.Flush()
}

func (d *TSDB) TextReport() string {
	b, err := json.Marshal(d.JsonReport())
	if err != nil {
		return fmt.Sprintf("%#v", *d)
	} else {
		return string(b)
	}
}

func (d *TSDB) JsonReport() interface{} {
	return map[string]interface{}{
		"foo": "bar",
	}
}

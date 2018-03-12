package reporter

import (
	"context"
	"fmt"
	"time"

	dlog "github.com/dyweb/gommon/log"
	"github.com/xephonhq/xephon-b/pkg/config"
	"github.com/xephonhq/xephon-b/pkg/metrics"
)

var _ Sink = (*Counter)(nil)

// Counter simply counts requests
type Counter struct {
	cfg config.CounterReporterConfig
	// TODO: failed requests
	totalRequests  int
	failedRequests int
	errorMessages  map[string]int
	codes          map[int]int

	// TODO: it seems impossible to count those based on information returned by worker, can just infer based on config?
	points int
	series int

	requestsSize int
	dataSize     int

	startTime time.Time
	endTime   time.Time

	log *dlog.Logger
}

func NewCounter(cfg config.CounterReporterConfig) *Counter {
	c := &Counter{
		cfg:           cfg,
		errorMessages: make(map[string]int),
		codes:         make(map[int]int),
	}
	dlog.NewStructLogger(log, c)
	return c
}

func (c *Counter) Run(ctx context.Context, resCh <-chan metrics.Response) {
	c.startTime = time.Now()
	for {
		select {
		case <-ctx.Done():
			c.log.Info("counter reporter stopped by context")
			goto END
		case res, ok := <-resCh:
			if !ok {
				c.log.Info("counter reporter stopped by closed channel")
				goto END
			}
			c.Record(res)
		}
	}
END:
	c.endTime = time.Now()
}

func (c *Counter) Record(res metrics.Response) {
	c.totalRequests++
	if res.GetError() {
		c.failedRequests++
		// default value for non existence key is 0
		c.errorMessages[res.GetErrorMessage()] = c.errorMessages[res.GetErrorMessage()] + 1
	}
	c.codes[res.GetCode()] = c.codes[res.GetCode()] + 1
	c.requestsSize += res.GetRequestSize()
}

func (c *Counter) Finalize() error {
	c.log.Info("finalize counter reporter, nothing to do")
	return nil
}

func (c *Counter) Flush() error {
	c.log.Info("flush counter reporter, nothing to do")
	return nil
}

func (c *Counter) TextReport() string {
	return fmt.Sprintf("%#v", *c)
}

func (c *Counter) JsonReport() interface{} {
	return map[string]interface{}{
		"totalRequests": c.totalRequests,
	}
}

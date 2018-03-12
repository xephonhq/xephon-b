package reporter

import (
	"context"
	"time"

	"fmt"
	dlog "github.com/dyweb/gommon/log"
	"github.com/xephonhq/xephon-b/pkg/metrics"
)

// Counter simply counts requests
type Counter struct {
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

func NewCounter() *Counter {
	c := &Counter{
		errorMessages: make(map[string]int),
		codes:         make(map[int]int),
	}
	dlog.NewStructLogger(log, c)
	return c
}

func (c *Counter) Run(ctx context.Context, resCh <-chan metrics.Result) {
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
			c.Report(&res)
		}
	}
END:
	c.endTime = time.Now()
}

// TODO: change to record? etc.?
func (c *Counter) Report(res *metrics.Result) {
	c.totalRequests++
	if res.Error {
		c.failedRequests++
		// default value for non existence key is 0
		c.errorMessages[res.ErrorMessage] = c.errorMessages[res.ErrorMessage] + 1
	}
	c.codes[res.Code] = c.codes[res.Code] + 1
	c.requestsSize += res.RequestSize
}

func (c *Counter) Finalize() {
	// FIXME: real printer, write to json or database etc.
	fmt.Printf("%#v", *c)
}

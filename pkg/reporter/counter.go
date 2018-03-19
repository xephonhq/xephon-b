package reporter

import (
	"context"
	"fmt"
	"time"

	"encoding/json"
	dlog "github.com/dyweb/gommon/log"
	"github.com/xephonhq/xephon-b/pkg/config"
	"github.com/xephonhq/xephon-b/pkg/metrics"
)

var _ Sink = (*Counter)(nil)

// Counter simply counts requests
type Counter struct {
	cfg            config.CounterReporterConfig
	totalRequests  int64
	failedRequests int64
	errorMessages  map[string]int
	codes          map[int]int

	totalPoints int
	// TODO: need to rely on config to calculate series
	//series int
	payloadSize int
	rawSize     int
	rawMetaSize int

	minLatency int64
	maxLatency int64
	avgLatency int64

	startTime time.Time
	endTime   time.Time

	// calculated
	requestPerSecond int64
	pointsPerSecond  int64
	duration         time.Duration
	dataRatio        int // in percentage, i.e. 90 means 90% is data 10% is meta

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
	c.minLatency = 999999999999999
	c.maxLatency = 0
	c.log.Infof("counter reporter start %s", c.startTime)
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
	c.log.Infof("counter reporter stop %s duration %s", c.endTime, c.endTime.Sub(c.startTime))
}

func (c *Counter) Record(res metrics.Response) {
	c.totalRequests++
	if res.GetError() {
		c.failedRequests++
		// default value for non existence key is 0
		c.errorMessages[res.GetErrorMessage()] = c.errorMessages[res.GetErrorMessage()] + 1
	}
	c.codes[res.GetCode()] = c.codes[res.GetCode()] + 1
	//c.log.Infof("points %d", res.GetPoints())
	c.totalPoints += res.GetPoints()
	c.payloadSize += res.GetPayloadSize()
	c.rawSize += res.GetRawSize()
	c.rawMetaSize += res.GetRawMetaSize()
	latency := res.GetEndTime() - res.GetStartTime()
	if latency < c.minLatency {
		c.minLatency = latency
	}
	if latency > c.maxLatency {
		c.maxLatency = latency
	}
	//c.avgLatency = (c.avgLatency*(c.totalRequests-1) + latency) / c.totalRequests // https://en.wikipedia.org/wiki/Moving_average Cumulative moving average
	c.avgLatency = c.avgLatency + (latency-c.avgLatency)/c.totalRequests
}

func (c *Counter) Finalize() error {
	c.log.Info("finalize counter reporter, calculate throughput")
	c.duration = c.endTime.Sub(c.startTime)
	// TODO: if the duration is less than 1s, we would have divide by zero ...
	duration := int64(c.duration / time.Second)
	c.requestPerSecond = c.totalRequests / duration
	c.pointsPerSecond = int64(c.totalPoints) / duration
	c.dataRatio = (c.rawSize - c.rawMetaSize) * 100 / c.rawSize
	return nil
}

func (c *Counter) Flush() error {
	c.log.Info("flush counter reporter, nothing to do")
	return nil
}

func (c *Counter) TextReport() string {
	b, err := json.Marshal(c.JsonReport())
	if err != nil {
		return fmt.Sprintf("%#v", *c)
	} else {
		return string(b)
	}
}

func (c *Counter) JsonReport() interface{} {
	return map[string]interface{}{
		"duration":       c.duration,
		"totalRequests":  c.totalRequests,
		"failedRequests": c.failedRequests,
		"totalPoints":    c.totalPoints,
		"payloadSize":    c.payloadSize,
		"rawSize":        c.rawSize,
		"rawMetaSize":    c.rawMetaSize,
		"dataRatio":      c.dataRatio,
	}
}

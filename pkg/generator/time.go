package generator

import (
	"github.com/dyweb/gommon/errors"
	"time"
)

var _ TimeGenerator = (*FixIntervalTime)(nil)

type FixIntervalTime struct {
	start     int64
	step      int64
	cur       int64
	precision time.Duration
}

// A Duration represents the elapsed time between two instants as an int64 nanosecond count.
func NewFixIntervalTime(interval time.Duration, precision time.Duration) (*FixIntervalTime, error) {
	now := time.Now().UnixNano()
	var (
		step  int64
		start int64
	)
	switch precision {
	case time.Second:
		start = now / 1000 / 1000 / 1000
		step = int64(interval) / 1000 / 1000 / 1000
	case time.Millisecond:
		start = now / 1000 / 1000
		step = int64(interval) / 1000 / 1000
	case time.Nanosecond:
		start = now
		step = int64(interval)
	default:
		return nil, errors.Errorf("unsupported precision %s", precision)
	}
	return &FixIntervalTime{
		start: start,
		step:  step,
		cur:   start,
	}, nil
}

func (g *FixIntervalTime) NextTime() int64 {
	g.cur += g.step
	return g.cur
}

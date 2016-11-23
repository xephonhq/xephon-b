package time

import (
	"errors"
	"time"
)

var ErrEndOfTime = errors.New("EOT")

type FixedIntervalTimeGenerator struct {
	//Start    time.Time
	//End      time.Time
	//Interval time.Duration
	start    int64
	end      int64
	current  int64
	interval int64
}

func NewFixedIntervalTimeGenerator(s time.Time, e time.Time, i time.Duration) *FixedIntervalTimeGenerator {
	return &FixedIntervalTimeGenerator{
		start:    s.UnixNano(),
		end:      e.UnixNano(),
		current:  s.UnixNano(),
		interval: i.Nanoseconds(),
	}
}

func (g *FixedIntervalTimeGenerator) NextTimestamp() (int64, error) {
	t := g.current
	g.current += g.interval
	if g.current >= g.end {
		return t, ErrEndOfTime
	}
	return t, nil
}

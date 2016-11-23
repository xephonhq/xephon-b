package generator

import (
	gt "github.com/xephonhq/xephon-b/pkg/generator/time"
	gv "github.com/xephonhq/xephon-b/pkg/generator/value"
	"testing"
	"time"
)

func TestNewGenerator(t *testing.T) {
	t.Parallel()
	v := gv.NewConstantIntGenerator(1)
	t.Log(v.NextInt())
	start := time.Now()
	end := time.Now().Add(time.Minute)
	step := time.Duration(10 * time.Second)
	tg := gt.NewFixedIntervalTimeGenerator(start, end, step)
	for {
		ts, err := tg.NextTimestamp()
		if err == gt.ErrEndOfTime {
			break
		}
		t.Log(ts)
	}
}

package generator

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/xephonhq/xephon-b/pkg/common"
)

func TestSeriesWithGenerator(t *testing.T) {
	assert := assert.New(t)

	name := "cpu.idle"
	start := time.Now().UnixNano()
	end := time.Now().Add(time.Minute).UnixNano()
	step := time.Duration(10 * time.Second).Nanoseconds()
	V := 10086
	// t.Logf("start %v, end %v, step %v", start, end, step)

	g := NewConstantIntGenerator(start, end, step, V)

	s := SeriesWithIntPointGenerator{Generator: g}
	s.Series = common.Series{Name: name}
	assert.Equal(name, s.Name)
	p, _ := s.Generator.Next()
	assert.Equal(start, p.TimeNano)
}

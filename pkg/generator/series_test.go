package generator

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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

	//s := SeriesWithIntPointGenerator{Series: name, Generator: &g}
	s := SeriesWithIntPointGenerator{Generator: g}
	assert.Equal(name, s.Name)
	assert.Equal(start, s.Generator.Next())
}

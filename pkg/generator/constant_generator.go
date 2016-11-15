package generator

// Constant generators

import (
	"github.com/xephonhq/xephon-b/pkg/common"
)

// ConstantIntPointGenerator generate int point over a time range in fixed interval
// start is included
// end may not be included
// number of points = (end - start) / step + 1
type ConstantIntPointGenerator struct {
	start   int64
	end     int64
	step    int64
	current int64
	V       int
}

// NewConstantIntPointGenerator create a generator, see test for example usage
func NewConstantIntPointGenerator(start int64, end int64, step int64, V int) *ConstantIntPointGenerator {
	// TODO: check
	return &ConstantIntPointGenerator{
		start:   start,
		end:     end,
		step:    step,
		current: start,
		V:       V,
	}
}

// Next return a new int point
// TODO: return pointer or value, use buffer, pool etc
func (c *ConstantIntPointGenerator) Next() (common.IntPoint, error) {
	p := common.IntPoint{
		V:        c.V,
		TimeNano: c.current,
	}
	c.current += c.step
	if c.current >= c.end {
		return p, ErrEndOfPoints
	}
	return p, nil
}

// IsValid check if all the required values are set and valid
func (c *ConstantIntPointGenerator) IsValid() bool {
	// TODO: /w\
	return true
}

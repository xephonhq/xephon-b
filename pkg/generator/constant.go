package generator

// Constant generators

import (
	"github.com/xephonhq/xephon-b/pkg/common"
)

// ConstantIntGenerator generate int point over a time range
// start is included
// end may not be included
type ConstantIntGenerator struct {
	start   int64
	end     int64
	step    int64
	current int64
	V       int
}

// NewConstantIntGenerator create a generator, see test for example usage
func NewConstantIntGenerator(start int64, end int64, step int64, V int) *ConstantIntGenerator {
	// TODO: check
	return &ConstantIntGenerator{
		start:   start,
		end:     end,
		step:    step,
		current: start,
		V:       V,
	}
}

// Next return a new int point
// TODO: return pointer or value, use buffer, pool etc
func (c *ConstantIntGenerator) Next() (common.IntPoint, error) {
	p := common.IntPoint{
		V: c.V,
		T: c.current,
	}
	c.current += c.step
	if c.current >= c.end {
		return p, ErrEndOfPoints
	}
	return p, nil
}

// IsValid check if all the required values are set and valid
func (c *ConstantIntGenerator) IsValid() bool {
	// TODO: /w\
	return true
}

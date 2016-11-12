package generator

import (
	"github.com/xephonhq/xephon-b/pkg/common"
)

type ConstantIntGenerator struct {
	start   int64
	end     int64
	step    int64
	current int64
	V       int
}

// Next return a new int point
// TODO: return pointer or value, buffer etc
func (c *ConstantIntGenerator) Next() (p common.IntPoint, last bool) {
	c.current += c.step
	return common.IntPoint{V: c.V, T: c.current}, false
}

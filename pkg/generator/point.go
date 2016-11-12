package generator

// IntPoint has integer value and unix nano timestamp in int64
type IntPoint struct {
	V int
	T int64
}

// DoublePoint has double value and unix nano timestamp in int64
type DoublePoint struct {
	V float64
	T int64
}

// IntPointGenerator generate integer value
// TODO: may change to some interface that support range
// TODO: may support channel
type IntPointGenerator interface {
	Next() (p IntPoint, last bool)
	IsValid() bool
}

// DoublePointGenerator generate double value
type DoublePointGenerator interface {
	Next() (p DoublePoint, last bool)
	IsValid() bool
}

type ConstantIntGenerator struct {
	start   int64
	end     int64
	step    int64
	current int64
	V       int
}

// Next return a new int point
// TODO: return pointer or value, buffer etc
func (c *ConstantIntGenerator) Next() (p IntPoint, last bool) {
	c.current += c.step
	return IntPoint{V: c.V, T: c.current}, false
}

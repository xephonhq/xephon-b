package value

type ConstantValueGenerator struct {
	vInt    int
	vDouble float64
}

func NewConstantIntGenerator(v int) *ConstantValueGenerator {
	return &ConstantValueGenerator{
		vInt: v,
	}
}

func (g *ConstantValueGenerator) NextInt() int {
	return g.vInt
}

func (g *ConstantValueGenerator) NextDouble() float64 {
	return g.vDouble
}

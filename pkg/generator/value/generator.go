package value

type ValueGenerator interface {
	NextInt() int
	NextDouble() float64
}

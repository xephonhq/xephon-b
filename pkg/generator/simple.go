package generator

import (
	"math/rand"
	"time"
)

type Constant struct {
	i int
	d float64
}

func NewConstantInt(v int) *Constant {
	return &Constant{
		i: v,
		d: float64(v),
	}
}

func NewConstantDouble(v float64) *Constant {
	return &Constant{
		i: int(v),
		d: v,
	}
}

func (g *Constant) NextInt() int {
	return g.i
}

func (g *Constant) NextDouble() float64 {
	return g.d
}

type Random struct {
	// don't use global rand methods because it is mutex protected
	r *rand.Rand
}

func NewRandom() *Random {
	return &Random{
		r: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (g *Random) NextInt() int {
	// Intn returns, as an int, a non-negative pseudo-random number in [0,n).
	// Int returns a non-negative pseudo-random int.
	return g.r.Int()
}

func (g *Random) NextDouble() float64 {
	// Float64 returns, as a float64, a pseudo-random number in [0.0,1.0).
	return g.r.Float64()
}

type Counter struct {
}

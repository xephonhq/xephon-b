package generator

import (
	"math/rand"
	"time"
)

var _ ValueGenerator = (*Constant)(nil)

type Constant struct {
	i int
	d float64
}

func NewConstant(i int, d float64) *Constant {
	return &Constant{
		i: i,
		d: d,
	}
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

var _ ValueGenerator = (*Random)(nil)

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

type Sequential struct {
	start    int64
	interval int64
	cur      int64
}

func NewSequential(start int64, interval int64) *Sequential {
	return &Sequential{
		start:    start,
		interval: interval,
		cur:      start,
	}
}

func (g *Sequential) NextSeq() int64 {
	g.cur += g.interval
	return g.cur
}

//type Counter struct {
//}

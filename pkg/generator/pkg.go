package generator

import "time"

type TimeGenerator interface {
	NextTime() time.Time
}

type IntGenerator interface {
	NextInt() int
}

type DoubleGenerator interface {
	NextDouble() float64
}

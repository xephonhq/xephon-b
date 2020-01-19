package generator

import (
	dlog "github.com/dyweb/gommon/log"
	"github.com/libtsdb/libtsdb-go/tspb"
)

var logReg = dlog.NewRegistry()
var log = logReg.Logger()

type SeriesGenerator interface {
	NextSeries() tspb.EmptySeries
}

type TimeGenerator interface {
	NextTime() int64
}

type IntGenerator interface {
	NextInt() int
}

type DoubleGenerator interface {
	NextDouble() float64
}

type ValueGenerator interface {
	IntGenerator
	DoubleGenerator
}

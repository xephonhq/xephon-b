package generator

import (
	pb "github.com/libtsdb/libtsdb-go/libtsdb/libtsdbpb"
	"github.com/xephonhq/xephon-b/pkg/util/logutil"
)

var log = logutil.NewPackageLogger()

type SeriesGenerator interface {
	NextSeries() pb.EmptySeries
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

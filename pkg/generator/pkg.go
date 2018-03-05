package generator

import (
	"github.com/xephonhq/xephon-b/pkg/util/logutil"
)

var log = logutil.NewPackageLogger()

// TODO: time generator need precision, this should be from config ...
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

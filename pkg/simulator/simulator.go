package simulator

import (
	"github.com/xephonhq/xephon-b/pkg/common"
	"github.com/xephonhq/xephon-b/pkg/serialize"
	"io"
	"github.com/xephonhq/xephon-b/pkg/generator"
)

// Simulator contains multiple series and represent one type of source of time series data
type Simulator interface {
	Type() string
	Series() []*common.Series
	AddSeriesWithIntPointGenerator(*generator.SeriesWithIntPointGenerator)
	AddSeriesWithDoublePointGenerator(*generator.SeriesWithDoublePointGenerator)
	SetSerializer(serialize.Serializer)
	SetWriter(io.Writer)
}

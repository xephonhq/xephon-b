package simulator

import (
	"github.com/xephonhq/xephon-b/pkg/common"
	"github.com/xephonhq/xephon-b/pkg/serialize"
	"io"
)

// Simulator contains multiple series and represent one type of source of time series data
type Simulator interface {
	Type() string
	Series() []*common.Series
	SetSerializer(*serialize.Serializer)
	SetWriter(*io.Writer)
}

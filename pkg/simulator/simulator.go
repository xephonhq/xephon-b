package simulator

import (
	"io"

	"github.com/xephonhq/xephon-b/pkg/serialize"
)

// Simulator contains multiple series and represent one type of source of time series data
type Simulator interface {
	Type() string
	SetSerializer(serialize.Serializer)
	SetWriter(io.Writer)
	Start()
}

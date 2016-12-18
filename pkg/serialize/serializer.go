package serialize

import (
	"github.com/xephonhq/xephon-b/pkg/common"
	"github.com/xephonhq/xephon-b/pkg/util"
)

// Short name use in machine simulator package
var log = util.Logger.NewEntryWithPkg("x.serialize")

// Serializer transform point with series into underlying format
type Serializer interface {
	WriteInt(*common.IntPointWithSeries) ([]byte, error)
	WriteDouble(*common.DoublePointWithSeries) ([]byte, error)
	ReadInt(s []byte) (*common.IntPointWithSeries, error)
	ReadDouble(s []byte) (*common.DoublePointWithSeries, error)
}

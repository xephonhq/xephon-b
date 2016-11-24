package serialize

import (
	"github.com/Sirupsen/logrus"
	"github.com/xephonhq/xephon-b/pkg/common"
	"github.com/xephonhq/xephon-b/pkg/util"
)

// Short name use in machine simulator package
var log = util.Logger.WithFields(logrus.Fields{
	"pkg": "x.serialize",
})

// Serializer transform point with series into underlying format
type Serializer interface {
	WriteInt(*common.IntPointWithSeries) ([]byte, error)
	WriteDouble(*common.DoublePointWithSeries) ([]byte, error)
}

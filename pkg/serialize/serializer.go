package serialize

import "github.com/xephonhq/xephon-b/pkg/common"

// Serializer transform point with series into underlying format
type Serializer interface {
	WriteInt(*common.IntPointWithSeries) ([]byte, error)
	WriteDouble(*common.DoublePointWithSeries) ([]byte, error)
}

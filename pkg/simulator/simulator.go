package simulator

import (
	"github.com/xephonhq/xephon-b/pkg/common"
)

// Simulator contains multiple series and represent one type of source of time series data
type Simulator interface {
	Name() string
	Series() []common.Series
}

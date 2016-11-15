package generator

import (
	"github.com/xephonhq/xephon-b/pkg/common"
)

// SeriesWithIntPointGenerator generate int point with series data attached to it
type SeriesWithIntPointGenerator struct {
	common.Series
	Generator *IntPointGenerator
}

// SeriesWithDoublePointGenerator generate double point with series data attached to it
type SeriesWithDoublePointGenerator struct {
	common.Series
	Generator *DoublePointGenerator
}

package generator

import (
	"github.com/xephonhq/xephon-b/pkg/common"
)

// SeriesWithIntPointGenerator generate int point with series data attached to it
type SeriesWithIntPointGenerator struct {
	common.Series
	// NOTE: you can't use *IntPointerGenerator
	// But when you pass a pointer to the Generator property, it will accept it
	// http://openmymind.net/Things-I-Wish-Someone-Had-Told-Me-About-Go/
	Generator IntPointGenerator
}

// SeriesWithDoublePointGenerator generate double point with series data attached to it
type SeriesWithDoublePointGenerator struct {
	common.Series
	Generator DoublePointGenerator
}

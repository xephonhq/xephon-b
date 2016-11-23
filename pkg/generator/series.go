package generator

import (
	"github.com/xephonhq/xephon-b/pkg/common"
	"github.com/xephonhq/xephon-b/pkg/generator/value"
)

type SeriesWithValueGenerator struct {
	common.Series
	ValueGenerator value.ValueGenerator
}

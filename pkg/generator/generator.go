package generator

import (
	"github.com/xephonhq/xephon-b/pkg/common"
)


// IntPointGenerator generate integer value
// TODO: may change to some interface that support built-in range operator, or may just add a Channel?
// TODO: may support channel
type IntPointGenerator interface {
	Next() (p common.IntPoint, last bool)
	IsValid() bool
}

// DoublePointGenerator generate double value
type DoublePointGenerator interface {
	Next() (p common.DoublePoint, last bool)
	IsValid() bool
}


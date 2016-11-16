package serialize

import (
	"fmt"

	"github.com/xephonhq/xephon-b/pkg/common"
)

//DebugSerializer change point with its series to human readable string
type DebugSerializer struct {
}

func (d *DebugSerializer) WriteInt(p *common.IntPointWithSeries) ([]byte, error) {
	// TODO: should use bytes.Buffer and maybe pool
	s := fmt.Sprintf("%s %d %d", p.Series, p.V, p.TimeNano)
	return []byte(s), nil
}

func (d *DebugSerializer) WriteDouble(p *common.DoublePointWithSeries) ([]byte, error) {
	// TODO: should use bytes.Buffer and maybe pool
	s := fmt.Sprintf("%s %0.2f %d", p.Series, p.V, p.TimeNano)
	return []byte(s), nil
}

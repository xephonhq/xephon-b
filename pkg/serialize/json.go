package serialize

import (
	"encoding/json"
	"github.com/xephonhq/xephon-b/pkg/common"
)

type JsonSerializer struct {
}

func (j *JsonSerializer) WriteInt(p *common.IntPointWithSeries) ([]byte, error) {
	s, err := json.Marshal(p)
	return s, err
}

func (j *JsonSerializer) WriteDouble(p *common.DoublePointWithSeries) ([]byte, error) {
	s, err := json.Marshal(p)
	return s, err
}

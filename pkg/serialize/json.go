package serialize

import (
	"encoding/json"

	"github.com/xephonhq/xephon-b/pkg/common"
)

type JsonSerializer struct {
}

func (j *JsonSerializer) WriteInt(p *common.IntPointWithSeries) ([]byte, error) {
	s, err := json.Marshal(p)
	// TODO： don't know if this append is efficient
	return append(s, '\n'), err
}

func (j *JsonSerializer) WriteDouble(p *common.DoublePointWithSeries) ([]byte, error) {
	s, err := json.Marshal(p)
	// TODO： don't know if this append is efficient
	return append(s, '\n'), err
}

func (j *JsonSerializer) ReadInt(s []byte) (*common.IntPointWithSeries, error) {
	p := common.IntPointWithSeries{}
	err := json.Unmarshal(s, &p)
	if err != nil {
		log.Warn(err)
	}
	return &p, err
}

func (j *JsonSerializer) ReadDouble(s []byte) (*common.DoublePointWithSeries, error) {
	p := common.DoublePointWithSeries{}
	err := json.Unmarshal(s, &p)
	if err != nil {
		log.Warn(err)
	}
	return &p, err
}
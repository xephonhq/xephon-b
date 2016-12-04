package kairosdb

import (
	"testing"
	"time"

	"github.com/xephonhq/xephon-b/pkg/common"
	"github.com/xephonhq/xephon-b/pkg/tsdb"
)

func TestTSDBClientInterface(t *testing.T) {
	var _ tsdb.TSDBClient = (*KairosDBHTTPClient)(nil)
}

func TestTSDBPayloadInterface(t *testing.T){
	t.Parallel()
	var _ tsdb.TSDBPayload = (*KairosDBPayload)(nil)
}

func TestKairosDBPayload_AddIntPoint(t *testing.T) {
	s := common.NewSeries("cpu")
	s.AddTag("os", "ubuntu")
	s.AddTag("type", "idle")
	p := common.IntPoint{TimeNano: time.Now().UnixNano(), V: 1}
	sp := common.IntPointWithSeries{Series: s, IntPoint: p}
	payload := NewKairosDBPayload()
	payload.AddIntPoint(&sp)
	payload.AddIntPoint(&sp)
	b, _ := payload.Bytes()
	t.Log(string(b))
}

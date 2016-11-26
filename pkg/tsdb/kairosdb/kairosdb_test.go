package kairosdb

import (
	"testing"
	"time"

	"github.com/xephonhq/xephon-b/pkg/common"
)

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

package kairosdb

import (
	"testing"
	"time"

	"github.com/xephonhq/xephon-b/pkg/common"
	"github.com/xephonhq/xephon-b/pkg/tsdb"
	"github.com/xephonhq/xephon-b/pkg/tsdb/config"
)

func TestTSDBClientInterface(t *testing.T) {
	var _ tsdb.TSDBClient = (*KairosDBHTTPClient)(nil)
}

func TestTSDBPayloadInterface(t *testing.T) {
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

func TestKairosDBHTTPClient_Initialize(t *testing.T) {
	c := config.TSDBClientConfig{
		Host: config.TSDBHostConfig{
			Address: "localhost",
			Port:    8080,
			SSL:     false,
		},
	}
	client := KairosDBHTTPClient{Config: c}
	//client.Initialize()
	client.Config.ConcurrentConnection = 100
	client.Initialize()
}

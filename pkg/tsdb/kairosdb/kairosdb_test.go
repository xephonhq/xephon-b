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

func TestKairosDBHTTPClient_Initialize_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Log("should panic when concurrent is not set")
			t.Fail()
		}
	}()
	c := config.TSDBClientConfig{
		Host: config.TSDBHostConfig{
			Address: "localhost",
			Port:    8080,
			SSL:     false,
		},
	}
	client := KairosDBHTTPClient{Config: c}
	client.Initialize()
}

func TestKairosDBHTTPClient_Initialize(t *testing.T) {
	c := config.TSDBClientConfig{
		Host: config.TSDBHostConfig{
			Address: "localhost",
			Port:    8080,
			SSL:     false,
		},
		ConcurrentConnection: 100,
	}
	client := KairosDBHTTPClient{Config: c}
	client.Initialize()
}
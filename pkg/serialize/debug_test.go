package serialize

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/xephonhq/xephon-b/pkg/common"

	"time"
	"fmt"
)

func TestDebugSerializer(t *testing.T) {
	assert := assert.New(t)
	ds := DebugSerializer{}
	name := "cpu.idle"
	s := common.NewSeries(name)
	s.AddTag("os", "ubuntu")
	s.AddTag("arch", "amd64")
	//assert.Equal("cpu.idle:os=ubuntu,arch=amd64,", s.String())
	p := common.IntPointWithSeries{Series: s}
	p.V = 123
	ts := time.Now().UnixNano()
	p.TimeNano = ts
	o := fmt.Sprintf("cpu.idle:os=ubuntu,arch=amd64, %d %d", 123, ts)
	w, _ := ds.WriteInt(&p)
	assert.Equal(o, string(w))

	p2 := common.DoublePointWithSeries{Series: s}
	p2.V = 12.03
	p2.TimeNano = ts
	o =  fmt.Sprintf("cpu.idle:os=ubuntu,arch=amd64, %0.2f %d", 12.03, ts)
	w, _ = ds.WriteDouble(&p2)
	assert.Equal(o, string(w))
}

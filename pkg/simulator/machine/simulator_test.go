package machine

import (
	"os"
	"testing"

	"time"

	"github.com/stretchr/testify/assert"
	"github.com/xephonhq/xephon-b/pkg/config"
	"github.com/xephonhq/xephon-b/pkg/serialize"
)

func TestGenerateDefaultMachine(t *testing.T) {
	assert := assert.New(t)
	m := GenerateDefaultMachine()
	assert.Equal("default-1", m.Name)
	assert.Equal("ubuntu16.04", m.OS)
	m = GenerateDefaultMachine()
	assert.Equal("default-2", m.Name)
}

func TestMachineSimulator(t *testing.T) {
	t.Parallel()
	c := config.MachineSimulatorConfig{
		Start: time.Now(),
		End:   time.Now().Add(time.Duration(time.Minute)),
		Step:  time.Duration(10 * time.Second),
		Num:   3,
	}
	sm := NewMachineSimulator(c)
	t.Log(sm)
	sm.SetWriter(os.Stdout)
	s := serialize.DebugSerializer{}
	sm.SetSerializer(&s)
	sm.Start()
}

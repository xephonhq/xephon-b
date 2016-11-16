package simulator

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xephonhq/xephon-b/pkg/common"
	"github.com/xephonhq/xephon-b/pkg/generator"
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

func TestAddSeries(t *testing.T) {
	assert := assert.New(t)
	s := common.Series{Name: "demo"}
	s2 := common.Series{Name: "demo-2"}
	sg := generator.SeriesWithIntPointGenerator{Series: s}
	sg2 := generator.SeriesWithDoublePointGenerator{Series: s2}
	sm := MachineSimulator{}
	sm.AddSeriesWithIntPointGenerator(&sg)
	sm.AddSeriesWithDoublePointGenerator(&sg2)
	assert.Equal(2, len(sm.Series()))
}

func TestAddDefaultMachine(t *testing.T) {
	assert := assert.New(t)
	sm := MachineSimulator{}
	sm.AddDefaultMachine()
	allSeries := sm.Series()
	assert.Equal("cpu_core", allSeries[0].TagKeys[1])
}

func TestMachineSimulator(t *testing.T) {
	t.Parallel()
	sm := MachineSimulator{}
	sm.AddDefaultMachine()
	sm.SetWriter(os.Stdout)
	s := serialize.DebugSerializer{}
	sm.SetSerializer(&s)
	sm.Start()
}

package simulator

import (
	"github.com/xephonhq/xephon-b/pkg/common"
	"github.com/xephonhq/xephon-b/pkg/serialize"
	"github.com/xephonhq/xephon-b/pkg/generator"
	"io"
	"sync/atomic"
	"fmt"
)

var defaulMachineNumber int64 = 0

type Machine struct {
	Name string
	OS   string
	RAM  int
	CPU  int
	DISK int
}

type MachineSimulator struct {
	start                          int64
	end                            int64
	step                           int64
	series                         []*common.Series
	seriesWithIntPointGenerator    []*generator.SeriesWithIntPointGenerator
	seriesWithDoublePointGenerator []*generator.SeriesWithDoublePointGenerator
	serializer                     serialize.Serializer
	writer                         io.Writer
}

func (ms *MachineSimulator) Type() string {
	return "machine"
}

func (ms *MachineSimulator) Series() []*common.Series {
	return ms.series
}

func (ms *MachineSimulator) SetSerializer(s serialize.Serializer) {
	ms.serializer = s
}

func (ms *MachineSimulator) SetWriter(w io.Writer) {
	ms.writer = w
}

func (ms *MachineSimulator) AddSeriesWithIntPointGenerator(g *generator.SeriesWithIntPointGenerator) {

}


func (ms *MachineSimulator) AddSeriesWithDoublePointGenerator(*generator.SeriesWithDoublePointGenerator) {

}

func GenerateDefaultMachine() Machine {
	num := atomic.AddInt64(&defaulMachineNumber, 1)
	return Machine{
		Name: fmt.Sprintf("default-%d", num),
		OS: "ubuntu16.04",
		CPU: 2,
		RAM: 2048,
		DISK: 1024000,
	}
}

func (ms *MachineSimulator) AddDefaultMachine() {
	ms.AddMachine(GenerateDefaultMachine())
}

func (ms *MachineSimulator) AddMachine(m Machine) {
	// add one machine == add a bunch of series, the simulator does NOT keep the machine
	// TODO: a base series and clone it to add more tags
	baseSeries := common.Series{
		Name:"machine",
	}
	baseSeries.AddTag("os", m.OS)
	
}



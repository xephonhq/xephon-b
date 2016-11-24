package machine

import (
	"io"
	"sync"

	"github.com/xephonhq/xephon-b/pkg/common"
	"github.com/xephonhq/xephon-b/pkg/config"
	gt "github.com/xephonhq/xephon-b/pkg/generator/time"

	"github.com/Sirupsen/logrus"
	"github.com/xephonhq/xephon-b/pkg/serialize"
	"github.com/xephonhq/xephon-b/pkg/util"
)

// Short name use in machine simulator package
var log = util.Logger.WithFields(logrus.Fields{
	"pkg": "x.s.machine",
})

var defaultMachineNumber int64 = 0

type MachineSimulator struct {
	config     config.MachineSimulatorConfig
	machines   []*Machine
	serializer serialize.Serializer
	writer     io.Writer
}

func (ms *MachineSimulator) Type() string {
	return "machine"
}

func (ms *MachineSimulator) SetSerializer(s serialize.Serializer) {
	ms.serializer = s
}

func (ms *MachineSimulator) SetWriter(w io.Writer) {
	ms.writer = w
}

func (ms *MachineSimulator) Start() {
	log.Debug("machine simulator start")
	intPointChan := make(chan *common.IntPointWithSeries)
	// TODO: passing byte array may not be efficient, but leave it to later ...
	serializedIntPointChan := make(chan []byte)
	// TODO: Add double point support
	// doublePointChan := make(chan *common.DoublePointWithSeries)
	var wg sync.WaitGroup
	var mg sync.WaitGroup
	wg.Add(1)
	go func() {
		mg.Add(len(ms.machines))
		for _, machine := range ms.machines {
			go func(m *Machine) {
				tg := gt.NewFixedIntervalTimeGenerator(ms.config.Start, ms.config.End, ms.config.Step)
				for {
					timestamp, err := tg.NextTimestamp()
					if err != nil {
						if err != gt.ErrEndOfTime {
							log.Warn(err)
						}
						break
					}
					for _, series := range m.intSeries {
						p := common.IntPoint{V: series.ValueGenerator.NextInt(), TimeNano: timestamp}
						intPointChan <- &common.IntPointWithSeries{IntPoint: p, Series: &series.Series}
					}
				}
				mg.Done()
			}(machine)
		}
		mg.Wait()
		close(intPointChan)
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		for p := range intPointChan {
			sp, err := ms.serializer.WriteInt(p)
			if err != nil {
				continue
			}
			serializedIntPointChan <- sp
		}
		close(serializedIntPointChan)
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		for sp := range serializedIntPointChan {
			ms.writer.Write(sp)
		}
		wg.Done()
	}()
	wg.Wait()
	log.Debug("machine simulator end")
}

func NewMachineSimulator(c config.MachineSimulatorConfig) *MachineSimulator {
	ms := MachineSimulator{config: c}
	for i := 0; i < c.Num; i++ {
		m := GenerateDefaultMachine()
		ms.AddMachine(&m)
	}
	return &ms
}

func (ms *MachineSimulator) AddMachine(m *Machine) {
	ms.machines = append(ms.machines, m)
}

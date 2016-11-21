package machine

import (
	"sync"

	"github.com/xephonhq/xephon-b/pkg/common"
	"github.com/xephonhq/xephon-b/pkg/config"
	gt "github.com/xephonhq/xephon-b/pkg/generator/time"
)

func NewMachineSimulator(c config.MachineSimulatorConfig) *MachineSimulator {
	ms := MachineSimulator{config: c}
	for i := 0; i < c.Num; i++ {
		m := GenerateDefaultMachineNew()
		ms.AddMachineNew(&m)
	}
	return &ms
}

// A temp file before remove the old code
func (ms *MachineSimulator) AddMachineNew(m *Machine) {
	ms.machines = append(ms.machines, m)
}

func (ms *MachineSimulator) StartNew() {
	intPointChan := make(chan *common.IntPointWithSeries)
	// TODO: passing byte array may not be efficient, but leave it to later ...
	serializedIntPointChan := make(chan []byte)
	// TODO:
	// doublePointChan := make(chan *common.DoublePointWithSeries)
	var wg sync.WaitGroup
	var mg sync.WaitGroup
	//wg.Add(len(ms.Series()))
	//for _, sIntGen := range ms.seriesWithIntPointGenerator {
	//	go func(g generator.IntPointGenerator, s common.Series) {
	//		for {
	//			p, err := g.Next()
	//			if err == generator.ErrEndOfPoints {
	//				break
	//			}
	//			intPointChan <- &common.IntPointWithSeries{IntPoint: p, Series: &s}
	//		}
	//		close(intPointChan)
	//		wg.Done()
	//	}(sIntGen.Generator, sIntGen.Series)
	//}
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
							// TODO: log or exit
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
}

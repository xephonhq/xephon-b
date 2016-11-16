package simulator

import (
	"fmt"
	"io"
	"sync"
	"sync/atomic"
	"time"

	"github.com/xephonhq/xephon-b/pkg/common"
	"github.com/xephonhq/xephon-b/pkg/generator"
	"github.com/xephonhq/xephon-b/pkg/serialize"
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
	ms.seriesWithIntPointGenerator = append(ms.seriesWithIntPointGenerator, g)
	ms.series = append(ms.series, &g.Series)
}

func (ms *MachineSimulator) AddSeriesWithDoublePointGenerator(g *generator.SeriesWithDoublePointGenerator) {
	ms.seriesWithDoublePointGenerator = append(ms.seriesWithDoublePointGenerator, g)
	ms.series = append(ms.series, &g.Series)
}

func (ms *MachineSimulator) Start() {
	fmt.Println("started!")
	// TODO: check series are valid and writer is SetWriter

	// TODO: config the start and end
	// assume we have then configured when add them

	// all the int points
	intPointChan := make(chan *common.IntPointWithSeries)
	// intSPointChan := make(chan []byte)
	tChan := make(chan int)
	// doublePointChan := make(chan *common.DoublePointWithSeries)
	var wg sync.WaitGroup
	wg.Add(len(ms.Series()))
	for _, sIntGen := range ms.seriesWithIntPointGenerator {
		go func(g generator.IntPointGenerator, s common.Series) {
			for {
				p, err := g.Next()
				fmt.Println("point generated")
				if err == generator.ErrEndOfPoints {
					break
				}
				intPointChan <- &common.IntPointWithSeries{IntPoint: p, Series: &s}
			}
			close(intPointChan)
			wg.Done()
		}(sIntGen.Generator, sIntGen.Series)
	}
	wg.Add(1)
	go func() {
		for p := range intPointChan {
			ms.serializer.WriteInt(p)
			fmt.Println("serialized!")
			// if err != nil {
			// 	// intSPointChan <- sp
			// }

			tChan <- 1

		}
		// close(intSPointChan)
		close(tChan)
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		fmt.Println("start writer routine")
		// for sp := range intSPointChan {
		for t := range tChan {

			// FIXME: this part is never executed
			fmt.Printf("need to write it! %d \n", t)
			// ms.writer.Write(sp)
		}
		wg.Done()
	}()
	wg.Wait()
}

func GenerateDefaultMachine() Machine {
	num := atomic.AddInt64(&defaulMachineNumber, 1)
	return Machine{
		Name: fmt.Sprintf("default-%d", num),
		OS:   "ubuntu16.04",
		CPU:  2,
		RAM:  2048,
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
		Name: "machine",
	}
	baseSeries.AddTag("os", m.OS)
	for i := 1; i < m.CPU; i++ {
		s := baseSeries
		s.AddTag("cpu_core", string(i))
		start := time.Now().UnixNano()
		end := time.Now().Add(time.Minute).UnixNano()
		step := time.Duration(10 * time.Second).Nanoseconds()
		V := 30
		g := generator.NewConstantIntPointGenerator(start, end, step, V)

		// let's just assume we only have cpu usage and it is a constant int
		ms.AddSeriesWithIntPointGenerator(&generator.SeriesWithIntPointGenerator{Series: s, Generator: g})
	}
	// what about mem and disk and etc ... e... wait until other thing is finished
}

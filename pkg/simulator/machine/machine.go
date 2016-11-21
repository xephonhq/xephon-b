package machine

import (
	"fmt"
	"sync/atomic"

	"github.com/xephonhq/xephon-b/pkg/common"
	"github.com/xephonhq/xephon-b/pkg/generator"
	"github.com/xephonhq/xephon-b/pkg/generator/value"
)

type Machine struct {
	Name         string
	OS           string
	RAM          int
	CPU          int
	DISK         int
	intSeries    []*generator.SeriesWithValueGenerator
	doubleSeries []*generator.SeriesWithValueGenerator
}

func GenerateDefaultMachine() Machine {
	num := atomic.AddInt64(&defaultMachineNumber, 1)
	return Machine{
		Name: fmt.Sprintf("default-%d", num),
		OS:   "ubuntu16.04",
		CPU:  2,
		RAM:  2048,
		DISK: 1024000,
	}
}

func GenerateDefaultMachineNew() Machine {
	num := atomic.AddInt64(&defaultMachineNumber, 1)
	m := Machine{
		Name: fmt.Sprintf("default-new-%d", num),
		OS:   "ubuntu16.04",
		CPU:  2,
		RAM:  2048,
		DISK: 1024000,
	}
	baseSeries := common.NewSeries("machine")
	for i := 0; i < 4; i++ {
		s := generator.SeriesWithValueGenerator{
			Series:         *baseSeries,
			ValueGenerator: value.NewConstantIntGenerator(1),
		}
		s.AddTag("host", m.Name)
		s.AddTag("os", m.OS)
		s.AddTag("cpu_core", fmt.Sprintf("%d", i))
		m.intSeries = append(m.intSeries, &s)
	}
	return m
}

package machine

import (
	"fmt"
	"sync/atomic"
)

type Machine struct {
	Name string
	OS   string
	RAM  int
	CPU  int
	DISK int
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

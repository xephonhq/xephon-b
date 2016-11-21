package machine

import "github.com/xephonhq/xephon-b/pkg/config"

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

}

package simulator

type Machine struct {
	OS   string
	RAM  int
	CPU  int
	DISK int
}

type MachineSimulator struct {
	start int64
	end   int64
	step  int64
}

func (ms *MachineSimulator) AddMachine(m Machine) {
	// add one machine == add a bunch of series
}

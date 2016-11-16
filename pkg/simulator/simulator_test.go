package simulator

import (
	"testing"
)

func TestSimulatorInterface(t *testing.T){
	t.Parallel()
	var _ Simulator = (*MachineSimulator)(nil)
}

package simulator

import (
	"github.com/xephonhq/xephon-b/pkg/simulator/machine"
	"testing"
)

func TestSimulatorInterface(t *testing.T) {
	t.Parallel()
	var _ Simulator = (*machine.MachineSimulator)(nil)
}

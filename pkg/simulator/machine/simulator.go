package machine

import (
	"io"

	"github.com/xephonhq/xephon-b/pkg/config"
	"github.com/xephonhq/xephon-b/pkg/serialize"
)

var defaultMachineNumber int64 = 0

type MachineSimulator struct {
	start      int64
	end        int64
	step       int64
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

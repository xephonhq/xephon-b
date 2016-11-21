package config

import (
	"testing"
	"github.com/xephonhq/xephon-b/pkg/util"
)

func TestReadMachineSimulatorConfigFromViper(t *testing.T) {
	util.ViperReadTestConfig()
	c := ReadMachineSimulatorConfigFromViper()
	t.Log(c)
}

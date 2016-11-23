package config

import (
	"github.com/xephonhq/xephon-b/pkg/util"
	"testing"
)

func TestReadMachineSimulatorConfigFromViper(t *testing.T) {
	util.ViperReadTestConfig()
	c := ReadMachineSimulatorConfigFromViper()
	t.Log(c)
}

package config

import (
	"time"

	"github.com/spf13/viper"
)

// MachineSimulatorConfig defines the configurable property for the machine simulator
type MachineSimulatorConfig struct {
	Num   int
	Start time.Time
	End   time.Time
	Step  int
}

// ReadMachineSimulatorConfigFromViper return a config struct using configuration in yml
func ReadMachineSimulatorConfigFromViper() *MachineSimulatorConfig {
	c := &MachineSimulatorConfig{}
	c.Num = viper.GetInt("simulator.machine.num")
	c.Start = viper.GetTime("simulator.machine.start")
	c.End = viper.GetTime("simulator.machine.end")
	c.Step = viper.GetInt("simulator.machine.step")
	return c
}

//func (config *MachineSimulatorConfig) ToMachineSimulator() *simulator.MachineSimulator {
//	sm := &simulator.MachineSimulator{
//		start:
//	}
//	return sm
//}

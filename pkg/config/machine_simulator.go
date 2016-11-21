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
	Step  time.Duration
}

// ReadMachineSimulatorConfigFromViper return a config struct using configuration in yml
func ReadMachineSimulatorConfigFromViper() *MachineSimulatorConfig {
	c := &MachineSimulatorConfig{}
	c.Num = viper.GetInt("simulator.machine.num")
	c.Start = viper.GetTime("simulator.machine.start")
	c.End = viper.GetTime("simulator.machine.end")
	// TODO: may remove the outer time.Duration?
	c.Step = time.Duration(time.Duration(viper.GetInt("simulator.machine.step")) * time.Second)
	return c
}

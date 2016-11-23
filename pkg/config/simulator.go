package config

var SupportedType = []string{"machine"}
var SupportedEncoding = []string{"debug", "json"}
var SupportedOutput = []string{"stdout", "file"}

type SimulatorConfig struct {
	Type           string
	Encoding       string
	Output         string
	OutputLocation string
}

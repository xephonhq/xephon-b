package common

// Database Names
const (
	KairosDB = "kairosdb"
	OpenTSDB = "opentsdb"
	InfluxDB = "influxdb"
)

// Default HTTP API Ports
const (
	KairosDBPort = 8080
	OpenTSDBPort = 4242
	InfuxDBPort  = 8086
)

// DefaultHTTPPorts is a map of default port numbers
var DefaultHTTPPorts = map[string]int{
	KairosDB: KairosDBPort,
	OpenTSDB: OpenTSDBPort,
	InfluxDB: InfuxDBPort,
}

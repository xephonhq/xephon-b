package config

import (
	"time"

	"github.com/libtsdb/libtsdb-go/libtsdb/config"
)

type XephonBConfig struct {
	Limit     string           `yaml:"limit"`
	Duration  time.Duration    `yaml:"duration"`
	Worker    WorkerConfig     `yaml:"worker"`
	Workload  string           `yaml:"workload"`
	Workloads []WorkloadConfig `yaml:"workloads"`
	Database  string           `yaml:"database"`
	Databases []DatabaseConfig `yaml:"databases"`
}

type WorkerConfig struct {
	Num int `yaml:"num"`
}

type WorkloadConfig struct {
	Name   string       `yaml:"name"`
	Type   string       `yaml:"type"`
	Series SeriesConfig `yaml:"series"`
	Time   TimeConfig   `yaml:"time"`
	Value  ValueConfig  `yaml:"value"`
}

type SeriesConfig struct {
	Num           int           `yaml:"num"`
	Churn         bool          `yaml:"churn"`
	ChurnDuration time.Duration `yaml:"churnDuration"`
	NumTags       int           `yaml:"numTags"`
}

type TimeConfig struct {
	//Type string `yaml:"type"`
	Interval time.Duration `yaml:"interval"`
}

type ValueConfig struct {
	Generator string                        `yaml:"generator"`
	Constant  *ConstantValueGeneratorConfig `yaml:"constant"`
}

type ConstantValueGeneratorConfig struct {
	Int    int     `yaml:"int"`
	Double float64 `yaml:"double"`
}

type DatabaseConfig struct {
	Name     string                       `yaml:"name"`
	Type     string                       `yaml:"type"`
	Influxdb *config.InfluxdbClientConfig `yaml:"influxdb"`
	Kairosdb *config.KairosdbClientConfig `yaml:"kairosdb"`
	Graphite *config.GraphiteClientConfig `yaml:"graphite"`
}

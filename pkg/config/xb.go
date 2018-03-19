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
	Reporter  string           `yaml:"reporter"`
	Reporters []ReporterConfig `yaml:"reporters"`
}

type WorkerConfig struct {
	Num int `yaml:"num"`
}

type WorkloadConfig struct {
	Name   string       `yaml:"name"`
	Batch  BatchConfig  `yaml:"batch"`
	Series SeriesConfig `yaml:"series"`
	Time   TimeConfig   `yaml:"time"`
	Value  ValueConfig  `yaml:"value"`
}

type BatchConfig struct {
	Series int `yaml:"series"`
	Points int `yaml:"points"`
}

type SeriesConfig struct {
	Prefix              string        `yaml:"prefix"`
	Num                 int           `yaml:"num"`
	Churn               bool          `yaml:"churn"`
	ChurnDuration       time.Duration `yaml:"churnDuration"`
	NumTags             int           `yaml:"numTags"`
	GroupPointsBySeries bool          `yaml:"groupPointsBySeries"`
}

type TimeConfig struct {
	//Type string `yaml:"type"`
	Interval time.Duration `yaml:"interval"`
}

type ValueConfig struct {
	Generator string                        `yaml:"generator"`
	Constant  *ConstantValueGeneratorConfig `yaml:"constant"`
	Random    *RandomValueGeneratorConfig   `yaml:"random"`
}

type ConstantValueGeneratorConfig struct {
	Int    int     `yaml:"int"`
	Double float64 `yaml:"double"`
}

type RandomValueGeneratorConfig struct {
	Min float64 `yaml:"min"`
	Max float64 `yaml:"max"`
}

type DatabaseConfig struct {
	Name     string                       `yaml:"name"`
	Type     string                       `yaml:"type"`
	Akumuli  *config.AkumuliClientConfig  `yaml:"akumuli"`
	Graphite *config.GraphiteClientConfig `yaml:"graphite"`
	Influxdb *config.InfluxdbClientConfig `yaml:"influxdb"`
	Kairosdb *config.KairosdbClientConfig `yaml:"kairosdb"`
}

type CounterReporterConfig struct {
	Foo string `yaml:"foo"`
}

type TSDBReporterConfig struct {
	// Sample is 1 from n samples, will ignore everything if n < 1
	Sample   int            `yaml:"sample"`
	Database DatabaseConfig `yaml:"database"`
}

type ReporterConfig struct {
	Name    string                 `yaml:"name"`
	Type    string                 `yaml:"type"`
	Counter *CounterReporterConfig `yaml:"counter"`
	TSDB    *TSDBReporterConfig    `yaml:"tsdb"`
}

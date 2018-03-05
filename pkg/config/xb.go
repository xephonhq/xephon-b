package config

import (
	"time"

	"github.com/libtsdb/libtsdb-go/libtsdb/config"
)

type XephonBConfig struct {
	Workload  string           `yaml:"workload"`
	Workloads []WorkloadConfig `yaml:"workloads"`
	Limit     string           `yaml:"limit"`
	Duration  time.Duration    `yaml:"duration"`
	Worker    WorkerConfig     `yaml:"worker"`
	Database  string           `yaml:"database"`
	Databases []DatabaseConfig `yaml:"databases"`
}

type WorkerConfig struct {
	Num int `yaml:"num"`
}

type WorkloadConfig struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
	// TODO: constant, random etc.
}

type DatabaseConfig struct {
	Name     string                       `yaml:"name"`
	Type     string                       `yaml:"type"`
	Influxdb *config.InfluxdbClientConfig `yaml:"influxdb"`
	Kairosdb *config.KairosdbClientConfig `yaml:"kairosdb"`
	Graphite *config.GraphiteClientConfig `yaml:"graphite"`
}

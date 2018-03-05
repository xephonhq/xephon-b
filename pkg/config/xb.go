package config

import "github.com/libtsdb/libtsdb-go/libtsdb/config"

type XephonBConfig struct {
	Workload  string           `yaml:"workload"`
	Workloads []WorkloadConfig `yaml:"workloads"`
	Database  string           `yaml:"database"`
	Databases []DatabaseConfig `yaml:"databases"`
}

type WorkloadConfig struct {
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

package config

import "fmt"

type TSDBHostConfig struct {
	Address string
	Port    int
	SSL     bool
}

func (c TSDBHostConfig) HostURL() string {
	if c.SSL {
		return fmt.Sprintf("https://%s:%d", c.Address, c.Port)
	}
	return fmt.Sprintf("http://%s:%d", c.Address, c.Port)
}

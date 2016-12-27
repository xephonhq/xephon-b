package config

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xephonhq/xephon-b/pkg/tsdb/common"
)

type TSDBHostConfig struct {
	Address string
	Port    int
	SSL     bool
}

func (c TSDBHostConfig) HostURL() string {
	// TODO: should trim the extra http(s) if user pass it as address
	if c.SSL {
		return fmt.Sprintf("https://%s:%d", c.Address, c.Port)
	}
	return fmt.Sprintf("http://%s:%d", c.Address, c.Port)
}

func NewDefaultHost(tsdb string) (TSDBHostConfig, error) {
	c := TSDBHostConfig{
		Address: "localhost",
		SSL:     false,
	}
	port, ok := common.DefaultHTTPPorts[tsdb]
	if !ok {
		return c, errors.New(fmt.Sprintf("%s doest not have default http port hardcoded", tsdb))
	}
	c.Port = port
	return c, nil
}

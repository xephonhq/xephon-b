package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSeriesString(t *testing.T) {
	assert := assert.New(t)

	name := "cpu.idle"
	s := NewSeries(name)
	s.AddTag("os", "ubuntu")
	s.AddTag("arch", "amd64")
	assert.Contains(s.String(), "cpu.idle", "os=ubuntu", "arch=amd64")
}

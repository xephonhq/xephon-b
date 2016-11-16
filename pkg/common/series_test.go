package common

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSeriesString(t *testing.T) {
	assert := assert.New(t)

	name := "cpu.idle"
	s := NewSeries(name)
	s.AddTag("os", "ubuntu")
	s.AddTag("arch", "amd64")
	assert.Equal("cpu.idle:os=ubuntu,arch=amd64,", s.String())
	assert.Equal("cpu.idle:os=ubuntu,arch=amd64,", fmt.Sprintf("%v", s))

}

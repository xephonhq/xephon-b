package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSeries_String(t *testing.T) {
	assert := assert.New(t)

	name := "cpu.idle"
	s := NewSeries(name)
	s.AddTag("os", "ubuntu")
	s.AddTag("arch", "amd64")
	assert.Contains(s.String(), "cpu.idle", "os=ubuntu", "arch=amd64")
}

func TestSeries_SortedKeys(t *testing.T) {
	assert := assert.New(t)

	s := NewSeries("dummy")
	s.AddTag("a", "123")
	s.AddTag("b", "789")
	s.AddTag("1a", "789")
	assert.Equal([]string{"1a", "a", "b"}, s.SortedKeys())
}

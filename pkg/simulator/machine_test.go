package simulator

import (
	"testing"
	"github.com/stretchr/testify/assert"

)

func TestGenerateDefaultMachine(t *testing.T){
	assert := assert.New(t)
	m := GenerateDefaultMachine()
	assert.Equal("default-1", m.Name)
	assert.Equal("ubuntu16.04", m.OS)
	m = GenerateDefaultMachine()
	assert.Equal("default-2", m.Name)

}

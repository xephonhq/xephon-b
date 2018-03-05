package generator

import "testing"

func TestConstant_NextInt(t *testing.T) {
	c := NewConstantInt(1)
	t.Log(c.NextInt())
	t.Log(c.NextInt())
	t.Log(c.NextInt())
	t.Log(c.NextDouble())
	c = NewConstantDouble(12.03)
	t.Log(c.NextDouble())
	t.Log(c.NextDouble())
	t.Log(c.NextDouble())
}

func TestRandom_NextInt(t *testing.T) {
	r := NewRandom()
	t.Log(r.NextInt())
	t.Log(r.NextInt())
	t.Log(r.NextInt())
	t.Log(r.NextDouble())
	t.Log(r.NextDouble())
	t.Log(r.NextDouble())
}
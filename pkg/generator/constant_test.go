package generator

import (
	"testing"
	"time"
)

func TestConstantIntGenerator(t *testing.T) {
	start := time.Now().UnixNano()
	end := time.Now().Add(time.Minute).UnixNano()
	step := time.Duration(10 * time.Second).Nanoseconds()

	g := NewConstantIntGenerator(start, end, step, 1)
	c := 0
	for {
		_, err := g.Next()
		c++
		if err == ErrEndOfPoints {
			break
		}
	}
}

package generator

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/xephonhq/xephon-b/pkg/common"
)

func TestConstantIntGenerator(t *testing.T) {
	assert := assert.New(t)

	start := time.Now().UnixNano()
	end := time.Now().Add(time.Minute).UnixNano()
	step := time.Duration(10 * time.Second).Nanoseconds()
	V := 10086
	// t.Logf("start %v, end %v, step %v", start, end, step)

	g := NewConstantIntGenerator(start, end, step, V)
	var points []common.IntPoint
	for {
		p, err := g.Next()
		points = append(points, p)
		if err == ErrEndOfPoints {
			break
		}
	}

	// NOTE: the number of points is not accurate like the one in the test for arbitray input
	assert.Equal(7, len(points))

	assert.Equal(V, points[0].V)

	assert.Equal(start, points[0].T)
	// t.Logf("end %v, last point %v", end, points[len(points)-1].T)
	assert.True(end >= points[len(points)-1].T)
}

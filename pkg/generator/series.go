package generator

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/dyweb/gommon/errors"
	pb "github.com/libtsdb/libtsdb-go/libtsdb/libtsdbpb"

	"github.com/xephonhq/xephon-b/pkg/config"
)

var _ SeriesGenerator = (*GenericSeries)(nil)

// TODO: libtsdb does not have definition for series
type GenericSeries struct {
	cfg       config.SeriesConfig
	series    []pb.EmptySeries
	r         *rand.Rand
	i         int
	len       int
	churnTime time.Time
}

func NewGenericSeries(cfg config.SeriesConfig) (*GenericSeries, error) {
	if cfg.Num < 1 {
		return nil, errors.Errorf("invalid series num %d", cfg.Num)
	}
	if cfg.NumTags < 1 {
		return nil, errors.Errorf("invalid numTags %d", cfg.NumTags)
	}
	g := &GenericSeries{
		cfg: cfg,
		i:   0,
		r:   rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	g.gen()
	return g, nil
}

func (g *GenericSeries) gen() {
	var series []pb.EmptySeries
	r := g.r.Intn(1000)
	for i := 0; i < g.cfg.Num; i++ {
		s := pb.EmptySeries{Name: fmt.Sprintf("%s.series.%d.%d", g.cfg.Prefix, r, i)}
		for j := 0; j < g.cfg.NumTags; j++ {
			t := pb.Tag{
				K: fmt.Sprintf("key%d%d", r, j),
				V: fmt.Sprintf("val%d%d", r, j),
			}
			s.Tags = append(s.Tags, t)
		}
		series = append(series, s)
	}
	g.i = 0
	g.len = len(series)
	g.series = series
	if g.cfg.Churn {
		g.churnTime = time.Now().Add(g.cfg.ChurnDuration)
	}
}

// cur cycle around series
func (g *GenericSeries) cur() int {
	if g.i < g.len {
		t := g.i
		g.i++
		return t
	} else {
		g.i = 0
		return 0
	}
}

func (g *GenericSeries) NextSeries() pb.EmptySeries {
	// generate a set of new series to simulate series churn
	if g.cfg.Churn && time.Now().After(g.churnTime) {
		g.gen()
	}
	return g.series[g.cur()]
}

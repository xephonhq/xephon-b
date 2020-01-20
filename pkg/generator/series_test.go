package generator

import (
	"testing"
	"time"

	"github.com/xephonhq/xephon-b/pkg/config"
)

func TestGenericSeries_NextSeries(t *testing.T) {
	s, err := NewGenericSeries(
		config.SeriesConfig{
			Prefix:        "test",
			Num:           2,
			NumTags:       5,
			Churn:         true,
			ChurnDuration: 5 * time.Millisecond,
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	// should cycle
	t.Log(s.NextSeries().Name)
	t.Log(s.NextSeries().Name)
	t.Log(s.NextSeries().Name)
	// should churn
	time.Sleep(10 * time.Millisecond)
	t.Log(s.NextSeries().Name)
	t.Log(s.NextSeries().Name)
	t.Log(s.NextSeries().Name)
}

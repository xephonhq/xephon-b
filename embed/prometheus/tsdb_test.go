package prometheus_test

import (
	"fmt"
	"io/ioutil"
	"testing"
	"time"

	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/storage/tsdb"
	"github.com/prometheus/prometheus/tsdb/wal"
)

func TestPromInput(t *testing.T) {
	if err := New(""); err != nil {
		t.Fatal(err)
	}
}

func New(inputPath string) error {
	p, err := ioutil.TempDir("", "")
	if err != nil {
		return fmt.Errorf("error while creating temporary directory: %w", err)
	}
	_, err = tsdb.Open(p, nil, nil, &tsdb.Options{
		WALSegmentSize:    wal.DefaultSegmentSize, // 128 MB
		RetentionDuration: 99999 * 24 * 60 * 60 * model.Duration(time.Second),
		MinBlockDuration:  model.Duration(2 * time.Hour),
		MaxBlockDuration:  model.Duration(2 * time.Hour),
	})
	if err != nil {
		return err
	}
	return nil
}

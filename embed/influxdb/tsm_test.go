package influxdb_test

import (
	"context"
	"io/ioutil"
	"os"
	"testing"

	"github.com/influxdata/influxdb/cmd/influxd/launcher"
)

func TestPromInput(t *testing.T) {
	if err := New("input.txt"); err != nil {
		t.Fatal(err)
	}
}

// FIXME: this still starts http server ...
func New(inputPath string) error {
	l := launcher.NewTestLauncher()
	if err := l.Run(context.Background()); err != nil {
		return err
	}
	if err := l.Setup(); err != nil {
		return err
	}

	inputFile, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer inputFile.Close()
	data, err := ioutil.ReadAll(inputFile)
	if err != nil {
		return err
	}
	if err := l.WritePoints(string(data)); err != nil {
		return err
	}
	return nil
}

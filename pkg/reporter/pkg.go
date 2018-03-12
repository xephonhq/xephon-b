package reporter

import (
	"context"
	"github.com/xephonhq/xephon-b/pkg/metrics"
	"github.com/xephonhq/xephon-b/pkg/util/logutil"
)

var log = logutil.NewPackageLogger()

type Sink interface {
	// Run drain from response channel until it is closed or the context is canceled
	Run(ctx context.Context, resCh <-chan metrics.Response)
	// Record is not go routine safe, it is used for testing, and Run should call Record
	Record(res metrics.Response)
	// Finalize do the final calculation and called when manager know workers all finished
	Finalize() error
	// Flush send the data to underlying storage (i.e. TSDB) if any
	Flush() error
	// TextReport is a summary in plain text for printing on terminal
	TextReport() string
	// JsonReport is a summary in json for machine
	JsonReport() interface{}
}

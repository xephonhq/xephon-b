package logutil

import (
	"github.com/dyweb/gommon/log"
	gommonlog "github.com/dyweb/gommon/util/logutil"
	libtsdblog "github.com/libtsdb/libtsdb-go/libtsdb/util/logutil"
)

var Registry = log.NewLibraryLogger()

func NewPackageLogger() *log.Logger {
	l := log.NewPackageLoggerWithSkip(1)
	Registry.AddChild(l)
	return l
}

func init() {
	// gain control of important libraries, NOTE: there could be duplicate and cycle when various library is involved
	// thus gommon/log would keep track of visited logger when doing recursive version of SetLevel and SetHandler
	Registry.AddChild(gommonlog.Registry)
	Registry.AddChild(libtsdblog.Registry)
}

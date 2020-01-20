package runner

import (
	dlog "github.com/dyweb/gommon/log"
)

var logReg = dlog.NewRegistry()
var log = logReg.Logger()

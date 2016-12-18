package main

import (
	"os"

	"github.com/xephonhq/xephon-b/pkg/util"
)

var log = util.Logger.NewEntryWithPkg("tsdb-proxy")

func main() {
	if RootCmd.Execute() != nil {
		os.Exit(-1)
	}
}

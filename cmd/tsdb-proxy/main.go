package main

import (
	"os"

	"github.com/xephonhq/xephon-b/pkg/util"
)

var log = util.Logger.NewEntry()

func main() {
	if RootCmd.Execute() != nil {
		os.Exit(-1)
	}
}

func init() {
	log.AddField("pkg", "tsdb-proxy")
}

package main

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/xephonhq/xephon-b/pkg/util"
)

var log = util.Logger.WithFields(logrus.Fields{
	"pkg": "tsdb-proxy",
})

func main() {
	if RootCmd.Execute() != nil {
		os.Exit(-1)
	}
}

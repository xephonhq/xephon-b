package main

import (
	"fmt"
	"os"
	"runtime"

	icli "github.com/at15/go.ice/ice/cli"
	"github.com/xephonhq/xephon-b/pkg/util/logutil"
)

const (
	myname = "xab"
)

// TODO: we should not use lib log as top registry, but they are same struct, so it should work
var log = logutil.Registry

var (
	version   string
	commit    string
	buildTime string
	buildUser string
	goVersion = runtime.Version()
)

var buildInfo = icli.BuildInfo{Version: version, Commit: commit, BuildTime: buildTime, BuildUser: buildUser, GoVersion: goVersion}

func main() {
	cli := icli.New(
		icli.Name(myname),
		icli.Description("Xephon-B version of http benchmark"),
		icli.Version(buildInfo),
		icli.LogRegistry(log),
	)
	root := cli.Command()
	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

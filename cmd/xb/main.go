package main

import (
	"fmt"
	"os"
	"runtime"

	icli "github.com/at15/go.ice/ice/cli"
	goicelog "github.com/at15/go.ice/ice/util/logutil"

	"github.com/xephonhq/xephon-b/pkg/config"
	"github.com/xephonhq/xephon-b/pkg/util/logutil"
)

const (
	myname = "xb"
)

// FIXME: debug logging is not working ....
var log = logutil.Registry

var (
	version   string
	commit    string
	buildTime string
	buildUser string
	goVersion = runtime.Version()
)

var buildInfo = icli.BuildInfo{Version: version, Commit: commit, BuildTime: buildTime, BuildUser: buildUser, GoVersion: goVersion}

var cli *icli.Root
var cfg config.XephonBConfig

func main() {
	cli = icli.New(
		icli.Name(myname),
		icli.Description("Xephon-B Time Series Benchmark cli"),
		icli.Version(buildInfo),
		icli.LogRegistry(log),
	)
	root := cli.Command()
	root.AddCommand(runCmd)
	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func mustLoadConfig() {
	if err := cli.LoadConfigTo(&cfg); err != nil {
		log.Fatal(err)
	}
}

func init() {
	log.AddChild(goicelog.Registry)
}

package main

import (
	"fmt"
	"os"
	"runtime"

	icli "github.com/dyweb/go.ice/cli"
	dlog "github.com/dyweb/gommon/log"
	"github.com/xephonhq/xephon-b/pkg/config"
)

const (
	myname = "xb"
)

var logReg = dlog.NewRegistry()
var log = logReg.Logger()

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
	)
	root := cli.Command()
	root.AddCommand(runCmd)
	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func mustLoadConfig() {
	panic("load config not implemented")
	// TODO: go.ice should use UnmarshalStrict, and might just use direct load ...
	//if err := cli.LoadConfigTo(&cfg); err != nil {
	//	log.Fatal(err)
	//}
}

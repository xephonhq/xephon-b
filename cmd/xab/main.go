package main

import (
	"fmt"
	"os"
	"runtime"

	icli "github.com/at15/go.ice/ice/cli"
	"github.com/xephonhq/xephon-b/pkg/util/logutil"
	"github.com/spf13/cobra"
	"net/http"
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

var testSeverCmd = &cobra.Command{
	Use:   "test-server",
	Short: "start test server",
	Long:  "Start test http server on 8080 with /ping /empty",
	Run: func(cmd *cobra.Command, args []string) {
		mux := http.NewServeMux()
		mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("pong"))
		})
		addr := "localhost:8080"
		srv := http.Server{
			Addr: addr,
		}
		srv.Handler = mux
		log.Infof("listen on %s try %s/ping in browser", addr, addr)
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	},
}

func main() {
	cli := icli.New(
		icli.Name(myname),
		icli.Description("Xephon-B version of http benchmark"),
		icli.Version(buildInfo),
		icli.LogRegistry(log),
	)
	root := cli.Command()
	root.AddCommand(testSeverCmd)
	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

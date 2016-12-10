package main

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/xephonhq/xephon-b/pkg/tsdb/config"
	"github.com/xephonhq/xephon-b/pkg/tsdb/influxdb"
	"github.com/xephonhq/xephon-b/pkg/tsdb/kairosdb"
)

var (
	db   = ""
	host = ""
	port = -1
)

// PingCmd check if database is alive
// ping --db kairosdb --host localhost --port 8080
var PingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Ping database",
	Long:  `Call certain database API, i.e. version to see if database is alive`,
	Run: func(cmd *cobra.Command, args []string) {
		if db == "" || host == "" || port == -1 {
			log.Error("must provide db, host, port")
			return
		}
		c := config.TSDBClientConfig{
			Host: config.TSDBHostConfig{
				Address: host,
				Port:    port,
				SSL:     false,
			},
		}

		db = strings.ToLower(db)
		switch db {
		case "kairosdb":
			client := kairosdb.KairosDBHTTPClient{Config: c}
			if err := client.Ping(); err != nil {
				log.Error(err.Error())
			} else {
				log.Info("KairosDB is working")
			}
		case "influxdb":
			client := influxdb.InfluxDBClient{Config: c}
			if err := client.Ping(); err != nil {
				log.Error(err.Error())
			} else {
				log.Info("InfluxDB is working")
			}
		default:
			log.Errorf("unsupported database %s", db)
		}
	},
}

func init() {
	PingCmd.Flags().StringVar(&db, "db", "", "target database type")
	PingCmd.Flags().StringVar(&host, "host", "", "host address i.e. localhost")
	PingCmd.Flags().IntVar(&port, "port", -1, "host port")

	RootCmd.AddCommand(PingCmd)
}

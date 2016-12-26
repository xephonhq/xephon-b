package influxdb

import (
	"net/http"

	"github.com/pkg/errors"
	"github.com/xephonhq/xephon-b/pkg/tsdb"
	"github.com/xephonhq/xephon-b/pkg/tsdb/config"
	"github.com/xephonhq/xephon-b/pkg/util"
	"github.com/xephonhq/xephon-b/pkg/util/requests"
)

const influxDBVersionHeader = "X-Influxdb-Version"

type InfluxDBClient struct {
	Config config.TSDBClientConfig
}

// Short name use in InfluxDB client package
var log = util.Logger.NewEntryWithPkg("x.tsdb.influxdb")

// Ping use InfluxDB /ping API to check if InfluxDB is alive
func (client *InfluxDBClient) Ping() error {
	// https://docs.influxdata.com/influxdb/v1.1/tools/api/
	pingURL := client.Config.Host.HostURL() + "/ping"
	res, err := requests.Get(pingURL)
	if err != nil {
		return errors.Wrapf(err, "can't reach InfluxDB via %s", pingURL)
	}
	if res.Res.StatusCode != http.StatusNoContent {
		return errors.Wrapf(err, "wrong status code returned %d, body is %s", res.Res.StatusCode, res.Text)
	}
	log.Info("InfluxDB version is " + res.Res.Header.Get(influxDBVersionHeader))
	return nil
}

func (client *InfluxDBClient) Put(p tsdb.TSDBPayload) error {
	return nil
}

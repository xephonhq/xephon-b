package influxdb

import (
	"errors"
	"net/http"

	"io/ioutil"

	"github.com/xephonhq/xephon-b/pkg/tsdb"
	"github.com/xephonhq/xephon-b/pkg/tsdb/config"
	"github.com/xephonhq/xephon-b/pkg/util"
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
	res, err := http.Get(client.Config.Host.HostURL() + "/ping")
	if err != nil {
		log.Warn("can't get InfluxDB version")
		log.Debug(err.Error())
		return err
	}
	defer res.Body.Close()
	// InfluxDB use header
	resContent, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Warn("can't read response body")
		log.Debug(err.Error())
		return err
	}
	if res.StatusCode != http.StatusNoContent {
		err = errors.New(string(resContent))
		log.Warnf("wrong status code returned %d", res.StatusCode)
		log.Debug(err.Error())
		return err
	}
	log.Info("InfluxDB version is " + res.Header.Get(influxDBVersionHeader))
	return nil
}

func (client *InfluxDBClient) Put(p tsdb.TSDBPayload) error {
	return nil
}

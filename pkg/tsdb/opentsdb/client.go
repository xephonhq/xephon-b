package opentsdb

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/xephonhq/xephon-b/pkg/tsdb"
	"github.com/xephonhq/xephon-b/pkg/tsdb/config"
	"github.com/xephonhq/xephon-b/pkg/util"
)

// Short name use in OpenTSDB client package
var log = util.Logger.WithFields(logrus.Fields{
	"pkg": "x.tsdb.opentsdb",
})

type OpenTSDBHTTPClient struct {
	Config config.TSDBClientConfig
}

type OpenTSDBTelnetClient struct {
}

func (client *OpenTSDBHTTPClient) Ping() error {
	res, err := http.Get(client.Config.Host.HostURL() + "/api/version")
	if err != nil {
		log.Warn("can't get OpenTSDB version")
		log.Debug(err.Error())
		return err
	}
	defer res.Body.Close()
	resContent, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Warn("can't read response body")
		log.Debug(err.Error())
		return err
	}
	var resData map[string]string
	if err := json.Unmarshal(resContent, &resData); err != nil {
		log.Warn("can't parse json")
		log.Debug(err.Error())
		return err
	}
	log.Info("OpenTSDB version is " + resData["version"])
	return nil
}

func (client *OpenTSDBHTTPClient) Put(p tsdb.TSDBPayload) error {
	return nil
}

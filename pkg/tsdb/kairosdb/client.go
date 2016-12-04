package kairosdb

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/xephonhq/xephon-b/pkg/tsdb"
	"github.com/xephonhq/xephon-b/pkg/tsdb/config"

	"encoding/json"
	"io/ioutil"

	"github.com/xephonhq/xephon-b/pkg/util"
)

// Short name use in loader package
var log = util.Logger.WithFields(logrus.Fields{
	"pkg": "x.tsdb.kairosdb",
})

type KairosDBHTTPClient struct {
	Config config.TSDBClientConfig
}

type KairosDBTelnetClient struct {
}

// Ping use KairosDB version API to check if it alive
func (client *KairosDBHTTPClient) Ping() error {
	res, err := http.Get(client.Config.Host.HostURL() + "/api/v1/version")
	if err != nil {
		log.Warn("can't get kairosdb version")
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
	log.Info("KairosDB version is " + resData["version"])
	return nil
}

func (client *KairosDBHTTPClient) Put(p tsdb.TSDBPayload) error {
	// cast it to its own payload
	payload, ok := p.(*KairosDBPayload)
	if !ok {
		// TODO: the logic here is quite ... strange, fatal would exit the program, but what if
		// people want to continue? They should not, it's a problem of developer not using the right type
		log.Fatal("must pass KairosDBPayload to KairosDBClient")
		return nil
	}
	payload.Bytes()
	return nil
}

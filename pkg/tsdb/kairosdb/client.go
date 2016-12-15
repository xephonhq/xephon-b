package kairosdb

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/xephonhq/xephon-b/pkg/tsdb"
	"github.com/xephonhq/xephon-b/pkg/tsdb/config"

	"encoding/json"
	"io/ioutil"

	"github.com/xephonhq/xephon-b/pkg/util"
	"errors"
)

// Short name use in KairosdDB client package
var log = util.Logger.WithFields(logrus.Fields{
	"pkg": "x.tsdb.kairosdb",
})

type KairosDBHTTPClient struct {
	Config      config.TSDBClientConfig
	transport   *http.Transport
	httpClients []*http.Client
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

func (client *KairosDBHTTPClient) Initialize() error {
	if client.Config.ConcurrentConnection < 1 {
		// TODO: panic
		log.Fatal("concurrent connection must be larger than 1")
		// FIXME: this error is never returned
		return errors.New("concurrent connection must be larger than 1")
	}
	client.transport = &http.Transport{}
	// create clients based on concurrent connection
	for i := 0; i < client.Config.ConcurrentConnection; i++ {
		// TODO: should allocate a fixed size array and assign
		client.httpClients = append(client.httpClients,
			&http.Client{Transport: client.transport})
	}
	return nil
}

func (client *KairosDBHTTPClient) Put(p tsdb.TSDBPayload) error {
	// cast it to its own payload
	payload, ok := p.(*KairosDBPayload)
	if !ok {
		// NOTE: it's fatal because user should know what the type of payload and choose right client
		log.Fatal("must pass KairosDBPayload to KairosDBClient")
		return nil
	}
	payload.Bytes()
	return nil
}

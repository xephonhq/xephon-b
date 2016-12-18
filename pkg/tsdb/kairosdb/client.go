package kairosdb

import (
	"net/http"

	"github.com/xephonhq/xephon-b/pkg/tsdb"
	"github.com/xephonhq/xephon-b/pkg/tsdb/config"

	"encoding/json"
	"io/ioutil"

	"github.com/xephonhq/xephon-b/pkg/util"
)

// Short name use in KairosdDB client package
var log = util.Logger.NewEntry()

type KairosDBHTTPClient struct {
	Config      config.TSDBClientConfig
	transport   *http.Transport
	httpClients []*http.Client
	requestChan chan *http.Request // TODO: maybe a buffered channel
	putURL      string
}

type KairosDBTelnetClient struct {
}

func init() {
	log.AddField("pkg", "x.tsdb.kairosdb")
}

// Ping use KairosDB version API to check if it alive
// Ping does not require Initialize to be called
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

// Initialize creates a bunch of http clients
func (client *KairosDBHTTPClient) Initialize() error {
	if client.Config.ConcurrentConnection < 1 {
		log.Panic("concurrent connection must be larger thant 1")
	}

	concurrency := client.Config.ConcurrentConnection
	client.transport = &http.Transport{
		MaxIdleConnsPerHost: concurrency,
	}
	// create clients based on concurrent connection
	// all clients share one transport
	for i := 0; i < concurrency; i++ {
		// TODO: should allocate a fixed size array and assign
		client.httpClients = append(client.httpClients,
			&http.Client{Transport: client.transport})
	}
	client.requestChan = make(chan *http.Request)
	client.putURL = client.Config.Host.HostURL() + "/api/v1/datapoints"
	// TODO: start go routine for each client
	// TODO: external methods to shut down all go routines, (close the channel seems to be the best)
	// TODO: give each go routine and id for debug

	return nil
}

// Shutdown stops all go routine
func (client *KairosDBHTTPClient) Shutdown() {
	close(client.requestChan)
}

// Put sends payload using one of the many http clients
func (client *KairosDBHTTPClient) Put(p tsdb.TSDBPayload) error {
	// cast it to its own payload
	payload, ok := p.(*KairosDBPayload)
	if !ok {
		log.Panic("must pass KairosDBPayload to KairosDBClient")
	}

	payload.Bytes()
	return nil
}

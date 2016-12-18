package kairosdb

import (
	"io"
	"net/http"
	"sync"

	"github.com/xephonhq/xephon-b/pkg/tsdb"
	"github.com/xephonhq/xephon-b/pkg/tsdb/config"

	"encoding/json"
	"io/ioutil"

	"github.com/xephonhq/xephon-b/pkg/util"
)

// Short name use in KairosdDB client package
var log = util.Logger.NewEntryWithPkg("x.tsdb.kairosdb")

type KairosDBHTTPClient struct {
	Config       config.TSDBClientConfig
	transport    *http.Transport
	httpClients  []*http.Client
	requestChan  chan *http.Request // TODO: maybe a buffered channel
	putURL       string
	initializeWg sync.WaitGroup
	shutdownWg   sync.WaitGroup
}

type KairosDBTelnetClient struct {
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

// Initialize creates a bunch of http clients and waits for every goroutine to start
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
	client.initializeWg.Add(concurrency)
	client.shutdownWg.Add(concurrency)
	for i := 0; i < concurrency; i++ {
		// TODO: a separate function for this
		go func(i int) {
			log.Debugf("http client %d routine started", i)
			// log.Infof("http client %d routine started", i)
			httpClient := client.httpClients[i]
			client.initializeWg.Done()
			for req := range client.requestChan {
				res, err := httpClient.Do(req)
				if err != nil {
					log.Warn(err)
				} else {
					io.Copy(ioutil.Discard, res.Body)
					// TODO: I wrote a 'FIXME: now the request is canceled' comment in mini-impl/ab code
					res.Body.Close()
				}
			}
			log.Debugf("http client %d routine stopped", i)
			// log.Infof("http client %d routine stopped", i)
			client.shutdownWg.Done()
		}(i)
	}
	client.initializeWg.Wait()
	return nil
}

// Shutdown close the request channel and waits for all the goroutine to return
func (client *KairosDBHTTPClient) Shutdown() {
	close(client.requestChan)
	client.shutdownWg.Wait()
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

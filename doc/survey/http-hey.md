# Hey - HTTP Benchmark in Go

- repo: https://github.com/rakyll/hey
- I had a re-implementation of it https://github.com/at15/mini-impl/tree/master/ab

## Parameters

- number of requests, n
- currency level, c, each worker run n/c requests
- query per second, q, `throttle` in `runWorker`
- duration, z, ignore n if z is set
  - num is set to `math.MaxInt32` and `Stop` is called when time has reached

## Features

- use `net/http/httptrace` https://golang.org/pkg/net/http/httptrace/
- [ ] [Feature : Distributed load sending](https://github.com/rakyll/hey/issues/91) use mqtt, implemented by community

## Implementation

- use new transport for each new client, which is caused low concurrency https://github.com/rakyll/hey/issues/31
- use `x/net/http2` to let transport use http2

````go
func (b *Work) runWorkers() {
	var wg sync.WaitGroup
	wg.Add(b.C)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
			ServerName:         b.Request.Host,
		},
		MaxIdleConnsPerHost: min(b.C, maxIdleConn),
		DisableCompression:  b.DisableCompression,
		DisableKeepAlives:   b.DisableKeepAlives,
		Proxy:               http.ProxyURL(b.ProxyAddr),
	}
	if b.H2 {
		http2.ConfigureTransport(tr)
	} else {
		tr.TLSNextProto = make(map[string]func(string, *tls.Conn) http.RoundTripper)
	}
	client := &http.Client{Transport: tr, Timeout: time.Duration(b.Timeout) * time.Second}

	// Ignore the case where b.N % b.C != 0.
	for i := 0; i < b.C; i++ {
		go func() {
			b.runWorker(client, b.N/b.C)
			wg.Done()
		}()
	}
	wg.Wait()
}
````

````go
func (b *Work) runWorker(client *http.Client, n int) {
	var throttle <-chan time.Time
	if b.QPS > 0 {
		throttle = time.Tick(time.Duration(1e6/(b.QPS)) * time.Microsecond)
	}

	if b.DisableRedirects {
		client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
	}
	for i := 0; i < n; i++ {
		// Check if application is stopped. Do not send into a closed channel.
		select {
		case <-b.stopCh:
			return
		default:
			if b.QPS > 0 {
				<-throttle
			}
			b.makeRequest(client)
		}
	}
}
````
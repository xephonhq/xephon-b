package config

// TSDBClientConfig control the concurrency of client
type TSDBClientConfig struct {
	Host                 TSDBHostConfig
	ConcurrentConnection int
	QPSPerClient         int
	Timeout              int
	EnableTrace          int // or trace level
}

package tsdb

type TSDBClient interface {
	Ping() error
	Put(p TSDBPayload) error
}

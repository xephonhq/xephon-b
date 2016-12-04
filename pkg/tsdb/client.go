package tsdb

type TSDBClient interface {
	Put(p TSDBPayload) error
}

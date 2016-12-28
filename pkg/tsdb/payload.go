package tsdb

type TSDBPayload interface {
	// TODO: would it be more efficient to pass *[]byte, idk
	Bytes() ([]byte, error)
}

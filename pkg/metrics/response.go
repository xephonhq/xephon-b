package metrics

// TODO: we can have the HttpTrace in libtsdb-go implement this interface
type Response interface {
	// GetError specifies if this result is an error
	GetError() bool
	// GetErrorMessage is empty if Error is false
	GetErrorMessage() string
	// GetCode is response code from server, normally http status code
	GetCode() int
	// GetStartTime is when the request is started
	GetStartTime() int64
	// GetEndTime is when the request is finished, response is drained, error or not
	GetEndTime() int64
	// GetRequestSize is the size of the payload excluding header etc.
	GetRequestSize() int
	// TODO: what about meta? it is a big cost indeed
	// GetDataSize is the actually data point size
	GetDataSize() int
}

var _ Response = (*DefaultResponse)(nil)

type DefaultResponse struct {
	Error        bool
	ErrorMessage string
	Code         int
	StartTime    int64
	EndTime      int64
	RequestSize  int
	DataSize     int
}

func (r *DefaultResponse) GetError() bool {
	return r.Error
}

func (r *DefaultResponse) GetErrorMessage() string {
	return r.ErrorMessage
}

func (r *DefaultResponse) GetCode() int {
	return r.Code
}

func (r *DefaultResponse) GetStartTime() int64 {
	return r.StartTime
}

func (r *DefaultResponse) GetEndTime() int64 {
	return r.EndTime
}

func (r *DefaultResponse) GetRequestSize() int {
	return r.RequestSize
}

func (r *DefaultResponse) GetDataSize() int {
	return r.DataSize
}

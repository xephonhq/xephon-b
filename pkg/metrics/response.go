package metrics

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
	// GetPayloadSize is the size of the payload excluding header etc.
	GetPayloadSize() int
	// GetRawSize is the size in byte for meta and points written without serialization, see libtsdbpb sizer.go
	GetRawSize() int
	// GetRawMetaSize is the size in byte for meta data written without serialization, series name tags etc.
	GetRawMetaSize() int
}

var _ Response = (*DefaultResponse)(nil)

type DefaultResponse struct {
	Error        bool
	ErrorMessage string
	Code         int
	StartTime    int64
	EndTime      int64
	PayloadSize  int
	RawSize      int
	RawMetaSize  int
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

func (r *DefaultResponse) GetPayloadSize() int {
	return r.PayloadSize
}

func (r *DefaultResponse) GetRawSize() int {
	return r.RawSize
}

func (r *DefaultResponse) GetRawMetaSize() int {
	return r.RawMetaSize
}

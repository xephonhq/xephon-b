package metrics

import "github.com/libtsdb/libtsdb-go/libtsdb"

type Response interface {
	libtsdb.Trace
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

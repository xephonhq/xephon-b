// Package metrics defines client side benchmark metrics returned by worker after each request
// it is a standalone package to avoid cycle import
package metrics

import "time"

// TODO: rename result
type Result struct {
	// Error specifies if this result is an error
	Error bool
	// ErrorMessage is empty if Error is false
	ErrorMessage string
	// Code is response code from server, normally http status code
	Code int
	// Time is when the request is started
	Time int64
	// Duration is the total time spent in a round trip
	Duration time.Duration
	// RequestSize is the size of the payload excluding header etc.
	RequestSize int
	// DataSize is the actually data point size TODO: what about meta? it is a big cost indeed
	DataSize int
}

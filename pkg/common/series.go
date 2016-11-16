package common

// NOTE: our series only have one value for each timestamp, because multiple values can be expand to single value
// and I am not sure how InfluxDB and Druid implement multiple values @czheo
//
// A multiple value series
// cpu.0 system=ubuntu, arch=amd64, usage=0.062, idle=0.034, 1412312312
// Expand to single value series
// cpu.0.usage system=ubuntu, arch=amd64, 0.062, 1412312312
// cpu.0.idle  system=ubuntu, arch=amd64, 0.034, 1412312312

// Series is a time series
type Series struct {
	// TODO: string or []byte
	Name string            `json:"name"`
	Tags map[string]string `json:"tag"`
}

// SeriesWithIntPoint is a series with int value points
type SeriesWithIntPoint struct {
	Series
	// TODO: store two arrays, one for timestamp, one for value might be more efficient
	Points []*IntPoint
}

// SeriesWithDoublePoint is a series with double value points
type SeriesWithDoublePoint struct {
	Series
	Points []*DoublePoint
}

// AddTag adds a key value pair WITHOUT ANY checking for duplication
func (s *Series) AddTag(key string, val string) {
	s.Tags[key] = val
}

// https://nathanleclaire.com/blog/2014/08/09/dont-get-bitten-by-pointer-vs-non-pointer-method-receivers-in-golang/
// NOTE: must use non-pointer receiver in order to use %s in fmt
func (s Series) String() string {
	// NOTE: used for debug only
	// name:k1=v1,k2=v2
	// TODO: more efficient
	str := s.Name + ":"
	for k, v := range s.Tags {
		str += k + "=" + v + ","
	}
	return str
}

// NewSeries return a series with its tag map intitialized
func NewSeries(name string) *Series {
	return &Series{
		Name: name,
		Tags: make(map[string]string),
	}
}

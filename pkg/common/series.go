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
	Name string
	// TODO: string or []byte
	TagKeys   []string
	TagValues []string
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

//func HashCode() string {
//
//}

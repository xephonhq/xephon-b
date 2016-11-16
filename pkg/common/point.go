package common

// Point represent a data point in a series
// ie: cpu.usage <2016-11-12-12:41:33, 0.062>, <2016-11-12-12:41:34, 0.078>
// cpu.usage is a series and it has two points, which shows different usage at different time

// IntPoint has integer value and unix nano timestamp in int64
type IntPoint struct {
	V        int   `json:"v"`
	TimeNano int64 `json:"t"`
}

// IntPointWithSeries contains a point to its series
type IntPointWithSeries struct {
	IntPoint
	*Series
}

// DoublePoint has double value and unix nano timestamp in int64
type DoublePoint struct {
	V        float64 `json:"v"`
	TimeNano int64   `json:"t"`
}

// DoublePointWithSeries contains a pointer to its series
type DoublePointWithSeries struct {
	DoublePoint
	*Series
}

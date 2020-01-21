module github.com/xephonhq/xephon-b/embed/prometheus

go 1.13

require (
	github.com/prometheus/common v0.8.0
	github.com/prometheus/prometheus v2.15.2+incompatible
)

replace github.com/prometheus/prometheus => ../../../../prometheus/prometheus

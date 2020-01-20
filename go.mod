module github.com/xephonhq/xephon-b

go 1.13

require (
	github.com/dyweb/go.ice v0.0.3
	github.com/dyweb/gommon v0.0.13
	github.com/libtsdb/libtsdb-go v0.0.0-20180319021657-419c24436f34
	github.com/spf13/cobra v0.0.5
)

replace github.com/dyweb/go.ice => ../../dyweb/go.ice

replace github.com/libtsdb/libtsdb-go => ../../libtsdb/libtsdb-go

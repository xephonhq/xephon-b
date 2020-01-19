package tsdbutil

import (
	"github.com/dyweb/gommon/errors"

	"github.com/libtsdb/libtsdb-go/database"
	//"github.com/libtsdb/libtsdb-go/libtsdb/client/akumuliw"
	//"github.com/libtsdb/libtsdb-go/libtsdb/client/graphitew"
	//"github.com/libtsdb/libtsdb-go/libtsdb/client/influxdbw"
	//"github.com/libtsdb/libtsdb-go/libtsdb/client/kairosdbw"
	"github.com/xephonhq/xephon-b/pkg/config"
)

func CreateClient(cfg config.DatabaseConfig) (database.TracedWriteClient, error) {
	switch cfg.Type {
	//case "akumuli":
	//	if cfg.Akumuli == nil {
	//		return nil, errors.New("akumuli is selected but no config")
	//	}
	//	return akumuliw.New(*cfg.Akumuli)
	//case "graphite":
	//	if cfg.Graphite == nil {
	//		return nil, errors.New("graphite is selected but no config")
	//	}
	//	return graphitew.New(*cfg.Graphite)
	//case "influxdb":
	//	if cfg.Influxdb == nil {
	//		return nil, errors.New("influxdb is selected but no config")
	//	}
	//	return influxdbw.New(*cfg.Influxdb)
	//case "kairosdb":
	//	if cfg.Kairosdb == nil {
	//		return nil, errors.New("kairosdb is selected but no config")
	//	}
	//	if cfg.Kairosdb.Telnet {
	//		return kairosdbw.NewTcp(*cfg.Kairosdb)
	//	}
	//	return kairosdbw.NewHttp(*cfg.Kairosdb)
	default:
		return nil, errors.Errorf("unknown database %s", cfg.Type)
	}
}

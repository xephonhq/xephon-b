package loader

import (
	"bufio"
	"io"

	"github.com/xephonhq/xephon-b/pkg/config"
	"github.com/xephonhq/xephon-b/pkg/serialize"
	"github.com/xephonhq/xephon-b/pkg/util"
)

// Short name use in loader package
var log = util.Logger.NewEntryWithPkg("x.loader")

type Loader struct {
	config     config.LoaderConfig
	source     io.Reader // use bufio.Scanner to read by line
	serializer serialize.Serializer
}

func NewLoader(c config.LoaderConfig) *Loader {
	l := &Loader{}
	l.source = c.Source
	l.config = c
	l.serializer = &serialize.JsonSerializer{}
	return l
}

func (l *Loader) Start() {
	scanner := bufio.NewScanner(l.source)
	for scanner.Scan() {
		sp, err := l.serializer.ReadInt(scanner.Bytes())
		if err != nil {
			log.Warn(err)
		}
		// TODO: this might be too much ouput for debug
		log.Debug(sp)
		// TODO: use channel to give this to client
	}
	if err := scanner.Err(); err != nil {
		log.Warn(err.Error())
	}
}

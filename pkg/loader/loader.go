package loader

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/xephonhq/xephon-b/pkg/config"
	"github.com/xephonhq/xephon-b/pkg/serialize"
	"github.com/xephonhq/xephon-b/pkg/util"
)

// Short name use in loader package
var log = util.Logger.WithFields(logrus.Fields{
	"pkg": "x.loader",
})

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
		//fmt.Println(scanner.Text()) // Println will add back the final '\n'
		sp, err := l.serializer.ReadInt(scanner.Bytes())
		if err != nil {
			log.Warn(err)
		}
		log.Debug(sp)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

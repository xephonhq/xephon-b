package generator

import "github.com/xephonhq/xephon-b/pkg/util"

var log = util.Logger.NewEntry()

func init() {
	log.AddField("pkg", "x.generator")
}

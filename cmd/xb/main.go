package main

import (
	"github.com/xephonhq/xephon-b/pkg/cmd"
)

func main() {
	// call me xb
	cmd.RootCmd.Use = "xb"
	cmd.Execute()
}

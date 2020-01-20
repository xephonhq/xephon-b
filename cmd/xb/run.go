package main

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/xephonhq/xephon-b/pkg/runner"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run benchmark",
	Run: func(cmd *cobra.Command, args []string) {
		mustLoadConfig()
		mgr, err := runner.NewManager(cfg)
		if err != nil {
			log.Fatal(err)
		}
		if err := mgr.Run(context.Background()); err != nil {
			log.Fatal(err)
		}
	},
}

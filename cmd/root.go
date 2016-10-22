package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// RootCmd is the top command, other commands should be its child
var RootCmd = &cobra.Command{
	Use:   "xephon-b",
	Short: "Time series benmark suite",
	Long:  `Xephon-B is a benmark suite with a micro benmark tool`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, Ayato")
	},
}

// Execute run the root command and return
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}

func init() {

}

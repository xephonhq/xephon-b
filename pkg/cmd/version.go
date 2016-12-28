package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Version need to be manually updated
const Version = "0.0.1-dev"

// VersionCmd print the version
var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show Xephon-B version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(Version)
	},
}

func init() {
	RootCmd.AddCommand(VersionCmd)
}

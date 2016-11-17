package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// VersionCmd print the version
var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show Xephon-B version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(Version)
	},
}

package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

var VersionCmd = &cobra.Command{
	Use: "version",
	Short: "Show Xephon-B version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(Version)
	},
}
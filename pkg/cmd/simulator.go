package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// SimulatorCmd generate and store simulated time series data
var SimulatorCmd = &cobra.Command{
	// TODO: alias is not shown in help
	Use:     "simulator",
	Aliases: []string{"sim"},
	Short:   "Simulate time series data for different scenario",
	Long:    "Simulate real word time series data, serialize and store to file",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("simulator hey!")
	},
}

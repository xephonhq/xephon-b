package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/xephonhq/xephon-b/pkg/config"
	"github.com/xephonhq/xephon-b/pkg/serialize"
	"github.com/xephonhq/xephon-b/pkg/simulator/machine"
)

// flags to bind
var (
	simulatorType           string
	simulatorDataEncoding   string
	simulatorOutput         string
	simulatorOutputLocation string
)

// SimulatorCmd generate and store simulated time series data
var SimulatorCmd = &cobra.Command{
	// TODO: alias is not shown in help
	Use:     "simulator",
	Aliases: []string{"sim"},
	Short:   "Simulate time series data for different scenario",
	Long:    "Simulate real word time series data, serialize and store to file",
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("triggered simulator command")
		// TODO: check if what user passed is not supported
		// TODO: simulator config is not even used
		if simulatorType != "machine" {
			log.Fatalf("we only support machine, no %s", simulatorType)
			return
		}

		c := config.ReadMachineSimulatorConfigFromViper()
		sm := machine.NewMachineSimulator(*c)
		switch simulatorDataEncoding {
		case "stdout":
			sm.SetWriter(os.Stdout)
		}
		switch simulatorType {
		case "json":
			sm.SetSerializer(&serialize.JsonSerializer{})
		case "debug":
			sm.SetSerializer(&serialize.DebugSerializer{})
		}
		sm.Start()
	},
}

func init() {
	SimulatorCmd.Flags().StringVar(&simulatorType, "type", "machine", "simluator type")
	SimulatorCmd.Flags().StringVar(&simulatorDataEncoding, "encoding", "debug", "serializer encoding")
	SimulatorCmd.Flags().StringVar(&simulatorOutput, "output", "stdout", "output type")
	SimulatorCmd.Flags().StringVar(&simulatorOutputLocation, "location", "give_me_a_name.dat", "output destination")
}

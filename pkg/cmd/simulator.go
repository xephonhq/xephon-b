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
		switch simulatorOutput {
		case "stdout":
			sm.SetWriter(os.Stdout)
		case "file":
			// try to open the file
			f, err := os.Create(simulatorOutputLocation)
			if err != nil {
				log.Error("can not create output file")
				log.Fatalf(err.Error())
				return
			}
			defer f.Close()
			sm.SetWriter(f)
		default:
			log.Fatalf("unsupported output type %s", simulatorOutput)
			return
		}
		switch simulatorDataEncoding {
		case "json":
			log.Debug("set encoding to json")
			sm.SetSerializer(&serialize.JsonSerializer{})
		case "debug":
			log.Debug("set encoding to debug")
			sm.SetSerializer(&serialize.DebugSerializer{})
		default:
			log.Fatalf("unsupported encoding %s", simulatorDataEncoding)
			return
		}
		sm.Start()
	},
}

func init() {
	SimulatorCmd.Flags().StringVar(&simulatorType, "type", "machine", "simluator type")
	SimulatorCmd.Flags().StringVar(&simulatorDataEncoding, "encoding", "debug", "serializer encoding")
	SimulatorCmd.Flags().StringVar(&simulatorOutput, "output", "stdout", "output type")
	SimulatorCmd.Flags().StringVar(&simulatorOutputLocation, "location", "give_me_a_name.dat", "output destination")

	RootCmd.AddCommand(SimulatorCmd)
}

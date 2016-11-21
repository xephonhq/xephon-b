package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xephonhq/xephon-b/pkg/config"
	"github.com/xephonhq/xephon-b/pkg/util"
)

// Version need to be manuaully updated
const Version = "0.0.1-dev"

// RootCmd is the top command, other commands should be its child
var RootCmd = &cobra.Command{
	Use:   "xephon-b",
	Short: "Time series benmark suite",
	Long:  `Xephon-B is a benmark suite with a micro benmark tool`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Xephon-B:" + Version + " Use `xb -h` for more information")
	},
}

// Execute run the root command and return
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&config.ConfigFile, "config", config.DefaultConfigFile, "config file (default is ./xephon-b.yml)")
	RootCmd.PersistentFlags().BoolVar(&config.Debug, "debug", false, "debug")

	RootCmd.AddCommand(VersionCmd)
	RootCmd.AddCommand(SimulatorCmd)
}

func initConfig() {
	if config.Debug {
		util.UseVerboseLog()
	}
	viper.AutomaticEnv()
	// TODO: check file existence
	viper.SetConfigFile(config.ConfigFile)
	viper.ReadInConfig()
}

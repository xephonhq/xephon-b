package cmd

import (
	"fmt"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xephonhq/xephon-b/pkg/config"
	"github.com/xephonhq/xephon-b/pkg/util"
)

// Short name use in machine simulator package
var log = util.Logger.WithFields(logrus.Fields{
	"pkg": "x.cmd",
})

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
}

func initConfig() {
	if config.Debug {
		util.UseVerboseLog()
	}
	viper.AutomaticEnv()
	// TODO: check file existence
	viper.SetConfigFile(config.ConfigFile)
	err := viper.ReadInConfig()
	if err != nil {
		log.Warn(err)
	} else {
		log.Debugf("config file %s is loaded", config.ConfigFile)
	}
}

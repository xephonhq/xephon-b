package main

import (
	"fmt"

	"github.com/dyweb/Ayi/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configFile        = ""
	defaultConfigFile = "tsdb-proxy.yml"
	debug             = false
)

var RootCmd = &cobra.Command{
	Use:   "tsdb-proxy",
	Short: "Time series database proxy",
	Long:  `TSDB proxy is a proxy for time series databases`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Proxy, hey!")
	},
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&configFile, "config", defaultConfigFile, "config file (default is ./xephon-b.yml)")
	RootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "debug")
}

func initConfig() {
	if debug {
		util.UseVerboseLog()
	}
	viper.AutomaticEnv()
	// TODO: check file existence
	viper.SetConfigFile(configFile)
	err := viper.ReadInConfig()
	if err != nil {
		log.Warn(err)
	} else {
		log.Debugf("config file %s is loaded", configFile)
	}
}

package util

// Test util

import (
	"path"
	"runtime"

	"github.com/spf13/viper"
)

// ViperReadTestConfig
func ViperReadTestConfig() {
	// Get current file path https://coderwall.com/p/_fmbug/go-get-path-to-current-file
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		log.Fatal("can't get current file path")
	}

	filePath := path.Join(path.Dir(filename), "../../xephon-b.yml")
	log.Debug(filePath)
	viper.SetConfigFile(filePath)
	viper.ReadInConfig()
}

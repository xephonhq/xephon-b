package util

// Test util

import (
	"os"
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

func TestKairosDB() bool {
	// TODO: maybe using testing.Short()
	if os.Getenv("TEST_KAIROSDB") == "1" {
		// TODO: maybe need to ping to make sure the db is running
		return true
	}
	return false
}

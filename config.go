package boot

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
)

const configPrefix string = "boot.config."

var defaultConfig map[string]interface{} = map[string]interface{}{
	"file":    "application.yaml",
	"path":    ".",
	"enabled": true,
}

func Log(message string) {
	fmt.Printf("[BOOT-plugin] %v |%s\n", time.Now().Format("2006-01-02 15:04:05"), message)
}

func InitDefaultConfig(prefix string, config map[string]interface{}) {
	for key, value := range config {
		viper.SetDefault(prefix+key, value)
	}
}

func initEnv() {
	// env start with BOOT_
	viper.SetEnvPrefix("BOOT")
	viper.AutomaticEnv()

}

func initDefault() {
	// default config file, if you don't set by envs and flags, use this
	InitDefaultConfig(configPrefix, defaultConfig)
}

func initLocalConfig() error {
	// if file contains suffix, then split
	file := viper.GetString(configPrefix + "file")
	viper.SetConfigFile(file)
	viper.AddConfigPath(viper.GetString(configPrefix + "path"))
	if err := viper.ReadInConfig(); err != nil {
		if e, ok := err.(*os.PathError); ok {
			Log(fmt.Sprintf("read config failed, %v", e))
		} else {
			return err
		}
	}
	return nil
}

func initConfig() error {
	initEnv()
	initDefault()
	if viper.GetBool(configPrefix + "enabled") {

		return initLocalConfig()
	}
	return nil
}

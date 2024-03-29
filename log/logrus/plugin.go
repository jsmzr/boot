package logrus

import (
	"fmt"
	"time"

	"github.com/jsmzr/boot"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type LogrusPlugin struct{}

const configPrefix = "boot.logging."

var defaultConfig map[string]interface{} = map[string]interface{}{
	"enabled":                true,
	"order":                  -5,
	"level":                  "INFO",
	"reportCaller":           false,
	"format.timestampFormat": time.RFC3339,
	"format.fullTimestamp":   true,
}

func (l *LogrusPlugin) Enabled() bool {
	return viper.GetBool(configPrefix + "enabled")
}
func (l *LogrusPlugin) Order() int {
	return viper.GetInt(configPrefix + "order")
}
func (l *LogrusPlugin) Load() error {
	level := viper.GetString(configPrefix + "level")
	boot.Log(fmt.Sprintf("logrus level is [%s]", level))
	if logLevel, err := logrus.ParseLevel(level); err != nil {
		return err
	} else {
		logrus.SetLevel(logLevel)
	}
	var format logrus.TextFormatter
	if err := viper.UnmarshalKey(configPrefix+"format", &format); err == nil {
		logrus.SetFormatter(&format)
	} else {
		boot.Log(fmt.Sprintf("logurs unmarshalKey [boot.logging.format] failed, %s", err.Error()))
	}
	logrus.SetReportCaller(viper.GetBool(configPrefix + "reportCaller"))
	return nil
}

func init() {
	boot.InitDefaultConfig(configPrefix, defaultConfig)
	boot.RegisterPlugin("logging", &LogrusPlugin{})
}

package logrus

import (
	"fmt"
	"strings"

	"github.com/jsmzr/boot"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type LogrusPlugin struct{}

const configPrefix = "boot.logging."

var defaultConfig map[string]interface{} = map[string]interface{}{
	"enabled":      true,
	"order":        -5,
	"level":        "INFO",
	"reportCaller": false,
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
	switch strings.ToUpper(level) {
	case "DEBUG":
		logrus.SetLevel(logrus.DebugLevel)
	case "WARN":
		logrus.SetLevel(logrus.WarnLevel)
	case "ERROR":
		logrus.SetLevel(logrus.ErrorLevel)
	default:
		logrus.SetLevel(logrus.InfoLevel)
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

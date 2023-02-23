package logrus

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func TestDefaultCOnfig(t *testing.T) {
	keyEnabled := "enabled"
	keyOrder := "order"
	keyLevel := "level"
	if defaultConfig[keyEnabled] != viper.GetBool(configPrefix+keyEnabled) {
		t.Fail()
	}

	if defaultConfig[keyOrder] != viper.GetInt(configPrefix+keyOrder) {
		t.Fail()
	}

	if defaultConfig[keyLevel] != viper.GetString(configPrefix+keyLevel) {
		t.Fail()
	}
}

func TestEnabled(t *testing.T) {
	logPlugin := LogrusPlugin{}
	keyEnabled := "enabled"
	viper.Set(configPrefix+keyEnabled, false)
	if logPlugin.Enabled() {
		t.Fail()
	}
	viper.Set(configPrefix+keyEnabled, true)
	if !logPlugin.Enabled() {
		t.Fail()
	}

}

func TestOrder(t *testing.T) {
	logPlugin := LogrusPlugin{}
	keyOrder := "order"
	viper.Set(configPrefix+keyOrder, 1)
	if logPlugin.Order() != 1 {
		t.Fail()
	}
}

func TestLoad(t *testing.T) {
	logPlugin := LogrusPlugin{}
	keyLevel := "level"
	// debug
	viper.Set(configPrefix+keyLevel, "debug")
	_ = logPlugin.Load()
	if logrus.DebugLevel != logrus.GetLevel() {
		t.Fail()
	}

	viper.Set(configPrefix+keyLevel, "DEBUG")
	_ = logPlugin.Load()
	if logrus.DebugLevel != logrus.GetLevel() {
		t.Fail()
	}

	viper.Set(configPrefix+keyLevel, "WARN")
	_ = logPlugin.Load()
	if logrus.WarnLevel != logrus.GetLevel() {
		t.Fail()
	}

	viper.Set(configPrefix+keyLevel, "ERROR")
	_ = logPlugin.Load()
	if logrus.ErrorLevel != logrus.GetLevel() {
		t.Fail()
	}

	viper.Set(configPrefix+keyLevel, "unknow")
	_ = logPlugin.Load()
	if logrus.InfoLevel != logrus.GetLevel() {
		t.Fail()
	}

	viper.Set(configPrefix+keyLevel, "INFO")
	_ = logPlugin.Load()
	if logrus.InfoLevel != logrus.GetLevel() {
		t.Fail()
	}
	// test failed format
	viper.Set(configPrefix+"format", "abc")
	_ = logPlugin.Load()
	viper.Set(configPrefix+"format", nil)
}

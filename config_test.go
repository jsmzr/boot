package boot

import (
	"testing"

	"github.com/spf13/viper"
)

func TestInitDefaultConfig(t *testing.T) {
	keyStr := "str"
	keyInt := "int"
	keyBool := "bool"
	testDefaultConfig := map[string]interface{}{
		keyStr:  "abc",
		keyInt:  123,
		keyBool: true,
	}
	testConfigPrefix := "demo.test."
	InitDefaultConfig(testConfigPrefix, testDefaultConfig)
	if testDefaultConfig[keyStr] != viper.GetString(testConfigPrefix+keyStr) {
		t.Fail()
	}

	if testDefaultConfig[keyInt] != viper.GetInt(testConfigPrefix+keyInt) {
		t.Fail()
	}

	if testDefaultConfig[keyBool] != viper.GetBool(testConfigPrefix+keyBool) {
		t.Fail()
	}
}

func TestInitDefault(t *testing.T) {
	_ = initConfig()
	if viper.GetString(configPrefix+"file") != defaultConfig["file"] {
		t.Fail()
	}

	if viper.GetString(configPrefix+"path") != defaultConfig["path"] {
		t.Fail()
	}
	if viper.GetBool(configPrefix+"enabled") != defaultConfig["enabled"] {
		t.Fail()
	}
}

func TestInitLocalConfigFailed(t *testing.T) {
	// load bad local file
	viper.Set("boot.config.file", "bad_config_test.yaml")
	if initLocalConfig() == nil {
		t.Fail()
	}
}
func TestInitLocalConfigFaile1(t *testing.T) {
	// not found file
	viper.Set("boot.config.file", "not_found.yaml")
	if initLocalConfig() != nil {
		t.Fail()
	}
}

func TestInitLocalConfig(t *testing.T) {
	viper.Set("boot.config.file", "application_test.yaml")
	if initLocalConfig() != nil {
		t.Fail()
	}

	if viper.GetString("test.str") != "foo" {
		t.Fail()
	}

	if viper.GetInt("test.int") != 2233 {
		t.Fail()
	}

	if viper.GetBool("test.bool") != false {
		t.Fail()
	}
}

func TestInitConfig(t *testing.T) {
	viper.Set("boot.config.file", "application_test.yaml")
	if initConfig() != nil {
		t.Fail()
	}
	viper.Set("boot.config.enabled", false)
	if initConfig() != nil {
		t.Fail()
	}
	// restore the default values
	viper.Set("boot.config.enabled", true)

}

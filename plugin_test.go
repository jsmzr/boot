package boot

import (
	"errors"
	"testing"

	"github.com/spf13/viper"
)

type TestPlugin struct{}

func (p *TestPlugin) Enabled() bool {
	return true
}

func (p *TestPlugin) Order() int {
	return 0
}

func (p *TestPlugin) Load() error {
	return nil
}

type TestUnavailable struct{}

func (p *TestUnavailable) Enabled() bool {
	return false
}

func (p *TestUnavailable) Order() int {
	return 1
}

func (p *TestUnavailable) Load() error {
	return nil
}

type TestFailedPlugin struct{}

func (p *TestFailedPlugin) Enabled() bool {
	return true
}

func (p *TestFailedPlugin) Order() int {
	return 2
}

func (p *TestFailedPlugin) Load() error {
	return errors.New("test failed plugin")
}

func TestRegisterPlugin(t *testing.T) {
	plugins = make(map[string]Plugin)
	name := "testPlugin"
	RegisterPlugin(name, &TestPlugin{})
}
func TestRegisterPluginFailed(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Failed()
		}
	}()
	plugins = make(map[string]Plugin)
	name := "testPlugin"
	RegisterPlugin(name, &TestPlugin{})
	RegisterPlugin(name, &TestPlugin{})

}

func TestPostProccessInitConfigFailed(t *testing.T) {
	viper.Set("boot.config.file", "bad_config_test.yaml")
	if PostProccess() == nil {
		t.Failed()
	}
}

func TestPostProccessPluginFailed(t *testing.T) {
	viper.Set("boot.config.file", "application_test.yaml")
	plugins = make(map[string]Plugin)
	RegisterPlugin("test", &TestPlugin{})
	RegisterPlugin("test1", &TestFailedPlugin{})
	if PostProccess() == nil {
		t.Fail()
	}
}
func TestPostProccess(t *testing.T) {
	viper.Set("boot.config.file", "application_test.yaml")
	plugins = make(map[string]Plugin)
	RegisterPlugin("test", &TestPlugin{})
	RegisterPlugin("test1", &TestUnavailable{})
	if PostProccess() != nil {
		t.Fail()
	}
}

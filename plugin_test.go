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

type TestInitializer struct{}

func (p *TestInitializer) Init() error {
	return nil
}

type TestInitializer1 struct{}

func (p *TestInitializer1) Init() error {
	return nil
}

type TestFailedInitializer struct{}

func (p *TestFailedInitializer) Init() error {
	return errors.New("test failed initializer")
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

func TestPostProcessInitConfigFailed(t *testing.T) {
	viper.Set("boot.config.file", "bad_config_test.yaml")
	if PostProcess() == nil {
		t.Failed()
	}
}

func TestPostProcessPluginFailed(t *testing.T) {
	viper.Set("boot.config.file", "application_test.yaml")
	plugins = make(map[string]Plugin)
	RegisterPlugin("test", &TestPlugin{})
	RegisterPlugin("test1", &TestFailedPlugin{})
	if PostProcess() == nil {
		t.Fail()
	}
}
func TestPostProcess(t *testing.T) {
	viper.Set("boot.config.file", "application_test.yaml")
	plugins = make(map[string]Plugin)
	RegisterPlugin("test", &TestPlugin{})
	RegisterPlugin("test1", &TestUnavailable{})
	if PostProcess() != nil {
		t.Fail()
	}
}

func TestPostProcessInitializer(t *testing.T) {
	initializers = make(map[string]Initializer)
	RegisterInitializer("test", &TestInitializer{})
	RegisterInitializer("test1", &TestInitializer1{})
	if PostProcess() != nil {
		t.Fail()
	}
}

func TestPostProcessInitializerFailed(t *testing.T) {
	initializers = make(map[string]Initializer)
	RegisterInitializer("test-failed", &TestFailedInitializer{})
	if PostProcess() == nil {
		t.Fail()
	}
}

func TestPostProcessTask(t *testing.T) {
	initializers = make(map[string]Initializer)
	tasks = make([]Task, 0)
	RegisterTask("test", "0 0/1 * * * *", func() {})
	if PostProcess() != nil {
		t.Fail()
	}
}

func TestPostProcessTaskFailed(t *testing.T) {
	initializers = make(map[string]Initializer)
	tasks = make([]Task, 0)
	RegisterTask("test", "", func() {})
	if PostProcess() == nil {
		t.Fail()
	}
}

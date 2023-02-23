package skywalking

import (
	"fmt"

	"github.com/SkyAPM/go2sky"
	"github.com/SkyAPM/go2sky/reporter"
	"github.com/jsmzr/boot"
	"github.com/spf13/viper"
)

type SkywalkingPlugin struct{}

const configPrefix = "boot.skywalking."

var defaultConfig map[string]interface{} = map[string]interface{}{
	"enabled": true,
	"order":   40,
	"name":    "boot-skywalking",
}

func (s *SkywalkingPlugin) Enabled() bool {
	return viper.GetBool(configPrefix + "enabled")
}

func (s *SkywalkingPlugin) Order() int {
	return viper.GetInt(configPrefix + "order")
}

func (s *SkywalkingPlugin) Load() error {
	address := viper.GetString(configPrefix + "address")
	if address == "" {
		return fmt.Errorf("skywalking backend address is null")
	}
	boot.Log(fmt.Sprintf("skywalking create grpc reporter by address: [%s]", address))
	r, err := reporter.NewGRPCReporter(address)
	if err != nil {
		return err
	}
	name := viper.GetString(configPrefix + "name")
	boot.Log(fmt.Sprintf("skywlaking report by name: [%s]", name))
	if t, err := go2sky.NewTracer(name, go2sky.WithReporter(r)); err != nil {
		return err
	} else {
		go2sky.SetGlobalTracer(t)
		return nil
	}
}

func init() {
	boot.InitDefaultConfig(configPrefix, defaultConfig)
	boot.RegisterPlugin("skywalking", &SkywalkingPlugin{})
}

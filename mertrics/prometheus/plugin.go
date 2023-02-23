package prometheus

import (
	"fmt"
	"net/http"

	"github.com/jsmzr/boot"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
)

type PrometheusPlugin struct{}

const configPrefix = "boot.prometheus."

var defaultConfig map[string]interface{} = map[string]interface{}{
	"enabled": true,
	"order":   30,
	"port":    9080,
	"path":    "/prometheus",
}

func (p *PrometheusPlugin) Enabled() bool {
	return viper.GetBool(configPrefix + "enabled")
}

func (p *PrometheusPlugin) Order() int {
	return viper.GetInt(configPrefix + "order")
}

func (p *PrometheusPlugin) Load() error {
	path := viper.GetString(configPrefix + "path")
	port := viper.GetInt(configPrefix + "port")
	boot.Log(fmt.Sprintf("prometheus start: [:%d%s]", port, path))
	go func() {
		mux := http.NewServeMux()
		mux.Handle(path, promhttp.Handler())
		if err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux); err != nil {
			boot.Log(fmt.Sprintf("prometheus start error:%s", err.Error()))
		}
	}()
	return nil
}

func init() {
	boot.InitDefaultConfig(configPrefix, defaultConfig)
	boot.RegisterPlugin("prometheus", &PrometheusPlugin{})
}

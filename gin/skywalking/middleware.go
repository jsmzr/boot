package skywalking

import (
	"github.com/SkyAPM/go2sky"
	v3 "github.com/SkyAPM/go2sky-plugins/gin/v3"
	"github.com/gin-gonic/gin"
	"github.com/jsmzr/boot"
	bootGin "github.com/jsmzr/boot/gin"
	"github.com/jsmzr/boot/tracer"
	"github.com/spf13/viper"
)

type SkywalkingMiddleware struct{}

const configPrefix = "boot.gin.middleware.skywalking."

var defaultConfig = map[string]interface{}{
	"enabled": true,
	"order":   10,
}

func (g *SkywalkingMiddleware) Load(e *gin.Engine) error {
	e.Use(v3.Middleware(e, go2sky.GetGlobalTracer()))
	e.Use(func(ctx *gin.Context) {
		tracer.ThreadLocal.Set(ctx.Request)
		ctx.Next()
		tracer.ThreadLocal.Remove()
	})
	return nil
}

func (g *SkywalkingMiddleware) Order() int {
	return viper.GetInt(configPrefix + "order")
}

func (g *SkywalkingMiddleware) Enabled() bool {
	return viper.GetBool(configPrefix + "enabled")
}

func init() {
	boot.InitDefaultConfig(configPrefix, defaultConfig)
	bootGin.RegisterMiddleware("skywalking", &SkywalkingMiddleware{})
}

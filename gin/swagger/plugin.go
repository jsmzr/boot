package swagger

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jsmzr/boot"
	bootGin "github.com/jsmzr/boot/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const configPrefix = "boot.swagger."

var defaultConfig map[string]interface{} = map[string]interface{}{
	"enabled": true,
	"order":   40,
	"path":    "/swagger/*any",
}

type GinSwaggerPlugin struct{}

func (g *GinSwaggerPlugin) Enabled() bool {
	return viper.GetBool(configPrefix + "enabled")
}

func (g *GinSwaggerPlugin) Order() int {
	return viper.GetInt(configPrefix + "order")
}

func (g *GinSwaggerPlugin) Load() error {
	path := viper.GetString(configPrefix + "path")
	boot.Log(fmt.Sprintf("swagger start by: [%s]", path))
	bootGin.RegisterRouter(func(e *gin.Engine) {
		e.GET(path, ginSwagger.WrapHandler(swaggerFiles.Handler))
	})
	return nil
}

func init() {
	boot.InitDefaultConfig(configPrefix, defaultConfig)
	boot.RegisterPlugin("swagger", &GinSwaggerPlugin{})
}

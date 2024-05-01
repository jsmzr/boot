package gin

import (
	"fmt"
	"sort"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jsmzr/boot"
	"github.com/spf13/viper"
)

type GinMiddleware interface {
	Load(*gin.Engine) error
	Order() int
	Enabled() bool
}

const configPrefix = "boot.gin."

var routerFunctions []func(*gin.Engine)
var middlewares = make(map[string]GinMiddleware)

var defaultConfig map[string]interface{} = map[string]interface{}{
	"port":                        8080,
	"mode":                        "debug",
	"middleware.logger.enabled":   true,
	"middleware.recovery.enabled": true,
}

func RegisterRouter(f func(*gin.Engine)) {
	routerFunctions = append(routerFunctions, f)
}

func RegisterMiddleware(name string, m GinMiddleware) {
	_, ok := middlewares[name]
	if ok {
		panic(fmt.Errorf("gin middleware [%s] already registerd", name))
	}
	log(fmt.Sprintf("Register [%s:%T] middleware", name, m))
	middlewares[name] = m
}

func log(message string) {
	fmt.Printf("[BOOT-GIN] %v| %s\n", time.Now().Format("2006-01-02 15:04:05"), message)
}

func initMiddleware(e *gin.Engine) error {
	if len(middlewares) == 0 {
		log("Not found Gin Middleware")
		return nil
	}
	values := make([]GinMiddleware, 0, len(middlewares))
	for _, v := range middlewares {
		if !v.Enabled() {
			log(fmt.Sprintf("[%T] middleware disable", v))
			continue
		}
		values = append(values, v)
	}
	sort.Slice(values, func(i, j int) bool {
		return values[i].Order() < values[j].Order()
	})
	for i := 0; i < len(values); i++ {
		if err := values[i].Load(e); err != nil {
			return err
		}
		log(fmt.Sprintf("Load [%T] middleware", values[i]))
	}

	return nil
}

func Run() error {
	// init base config
	boot.InitDefaultConfig(configPrefix, defaultConfig)
	// init plugin
	if err := boot.PostProcess(); err != nil {
		return err
	}
	gin.SetMode(viper.GetString(configPrefix + "mode"))
	// gin start
	e := gin.New()
	if viper.GetBool(configPrefix + "middleware.logger.enabled") {
		e.Use(gin.Logger())
	}
	if viper.GetBool(configPrefix + "middleware.recovery.enabled") {
		e.Use(gin.Recovery())
	}

	if err := initMiddleware(e); err != nil {
		return err
	}
	for _, f := range routerFunctions {
		f(e)
	}
	return e.Run(fmt.Sprintf(":%d", viper.GetInt(configPrefix+"port")))
}

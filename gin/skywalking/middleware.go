package skywalking

import (
	"github.com/SkyAPM/go2sky"
	v3 "github.com/SkyAPM/go2sky-plugins/gin/v3"
	"github.com/gin-gonic/gin"
	bootGin "github.com/jsmzr/boot/gin"
	"github.com/jsmzr/boot/tracer"
)

type GinSkywalkingMiddle struct{}

func (g *GinSkywalkingMiddle) Load(e *gin.Engine) error {
	e.Use(v3.Middleware(e, go2sky.GetGlobalTracer()))
	e.Use(func(ctx *gin.Context) {
		tracer.ThreadLocal.Set(ctx.Request)
		ctx.Next()
		tracer.ThreadLocal.Remove()
	})
	return nil
}

func (g *GinSkywalkingMiddle) Order() int {
	return 10
}

func init() {
	bootGin.RegisterMiddleware("skywalking", &GinSkywalkingMiddle{})
}

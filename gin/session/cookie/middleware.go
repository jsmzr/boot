package cookie

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	bootGin "github.com/jsmzr/boot/gin"
	common "github.com/jsmzr/boot/gin/session"
	"github.com/spf13/viper"
)

type CookieMiddleware struct{}

const configPrefix = "boot.gin.session."

func (c *CookieMiddleware) Load(e *gin.Engine) error {
	name := viper.GetString(configPrefix + "name")
	secret := viper.GetString(configPrefix + "secret")
	store := cookie.NewStore([]byte(secret))
	common.SetStoreOptions(store)
	e.Use(sessions.Sessions(name, store))
	return nil
}

func (c *CookieMiddleware) Enabled() bool {
	return viper.GetBool(configPrefix + "enabled")
}

func (c *CookieMiddleware) Order() int {
	return viper.GetInt(configPrefix + "order")
}

func init() {
	bootGin.RegisterMiddleware("session", &CookieMiddleware{})
}

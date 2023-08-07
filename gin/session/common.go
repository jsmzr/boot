package session

import (
	"github.com/gin-contrib/sessions"
	"github.com/jsmzr/boot"
	"github.com/spf13/viper"
)

const configPrefix = "boot.gin.session."

var defaultConfig = map[string]interface{}{
	"name":     "gsession",
	"path":     "/",
	"maxAge":   86400 * 30,
	"secure":   false,
	"httpOnly": false,
	// 1 ~ 4, see http/cookie.go
	"sameSite": 1,
	"enabled":  true,
	"order":    13,
}

func SetStoreOptions(store sessions.Store) {
	options := sessions.Options{
		Path:     viper.GetString(configPrefix + "path"),
		MaxAge:   viper.GetInt(configPrefix + "maxAge"),
		Secure:   viper.GetBool(configPrefix + "secure"),
		HttpOnly: viper.GetBool(configPrefix + "httpOnly"),
	}
	domain := viper.GetString(configPrefix + "domain")
	if domain != "" {
		options.Domain = domain
	}
	store.Options(options)
}

func init() {
	boot.InitDefaultConfig(configPrefix, defaultConfig)
}

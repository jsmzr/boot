package redigo

import (
	"errors"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	bp "github.com/jsmzr/boot/cache/redigo/pool"
	bootGin "github.com/jsmzr/boot/gin"
	common "github.com/jsmzr/boot/gin/session"
	"github.com/spf13/viper"
)

type RedisMiddleware struct{}

const configPrefix = "boot.gin.session."

func (r *RedisMiddleware) Load(e *gin.Engine) error {
	if bp.Instance == nil {
		return errors.New("not found redigo pool")
	}
	store, err := redis.NewStoreWithPool(bp.Instance, []byte(viper.GetString(configPrefix+"secret")))
	if err != nil {
		return err
	}
	common.SetStoreOptions(store)
	keyPrefix := viper.GetString(configPrefix + "keyPrefix")
	if keyPrefix != "" {
		if err := redis.SetKeyPrefix(store, keyPrefix); err != nil {
			return err
		}
	}
	e.Use(sessions.Sessions(viper.GetString(configPrefix+"name"), store))
	return nil
}

func (r *RedisMiddleware) Enabled() bool {
	return viper.GetBool(configPrefix + "enabled")
}

func (r *RedisMiddleware) Order() int {
	return viper.GetInt(configPrefix + "order")
}

func init() {
	bootGin.RegisterMiddleware("session", &RedisMiddleware{})
}

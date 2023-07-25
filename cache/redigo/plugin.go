package redigo

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/jsmzr/boot"
	bp "github.com/jsmzr/boot/cache/redigo/pool"
	"github.com/spf13/viper"
)

const configPrefix = "boot.redis."

var defaultConfig = map[string]interface{}{
	"enabled":              true,
	"order":                12,
	"host":                 "127.0.0.1",
	"port":                 6379,
	"database":             0,
	"readTimeout":          0,
	"writeTimeout":         0,
	"useTLS":               false,
	"skipVerify":           false,
	"pool.macActive":       10,
	"pool.maxIdle":         5,
	"pool.idleTimeout":     240,
	"pool.maxConnLifetime": 0,
}

type RedigoPlugin struct{}

func (r *RedigoPlugin) Load() error {
	timeout := time.Duration(viper.GetInt(configPrefix + "idleTimeout"))
	address := fmt.Sprintf("%s:%d", viper.GetString(configPrefix+"host"), viper.GetInt(configPrefix+"port"))
	pool := redis.Pool{
		MaxIdle:         viper.GetInt(configPrefix + "pool.maxIdle"),
		IdleTimeout:     timeout * time.Second,
		MaxActive:       viper.GetInt(configPrefix + "pool.maxActive"),
		MaxConnLifetime: time.Duration(viper.GetInt(configPrefix + "pool.maxConnLifetime")),
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", address,
				redis.DialDatabase(viper.GetInt(configPrefix+"database")),
				redis.DialPassword(viper.GetString(configPrefix+"password")),
				redis.DialReadTimeout(time.Duration(viper.GetInt(configPrefix+"readTimeout"))),
				redis.DialWriteTimeout(time.Duration(viper.GetInt(configPrefix+"writeTimeout"))),
				redis.DialUseTLS(viper.GetBool(configPrefix+"useTLS")),
				redis.DialTLSSkipVerify(viper.GetBool(configPrefix+"skipVerify")),
			)
		},
	}
	c := pool.Get()
	if data, err := c.Do("PING"); err != nil {
		return err
	} else {
		if data != "PONG" {
			return fmt.Errorf("ping return: %+v", data)
		}
	}
	bp.Instance = &pool
	return nil
}

func (r *RedigoPlugin) Enabled() bool {
	return viper.GetBool(configPrefix + "enabled")
}

func (r *RedigoPlugin) Order() int {
	return viper.GetInt(configPrefix + "order")
}

func init() {
	boot.InitDefaultConfig(configPrefix, defaultConfig)
	boot.RegisterPlugin("redigo", &RedigoPlugin{})
}

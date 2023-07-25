package pool

import "github.com/gomodule/redigo/redis"

var Instance *redis.Pool

func Get() redis.Conn {
	return Instance.Get()
}

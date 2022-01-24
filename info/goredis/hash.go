package goredis

import (
	"github.com/gomodule/redigo/redis"
)

func HMSET(key string, value interface{}) (err error) {
	_, err = conn().Do("HMSET", redis.Args{}.Add(key).AddFlat(value)...)

	return
}

func HGETALL(key string, out interface{}) (err error) {
	data, err := redis.Values(conn().Do("HGETALL", key))
	if err != nil {
		return
	}

	ScanStruct(data, out)

	return
}

func ScanStruct(data []interface{}, out interface{}) (err error) {
	err = redis.ScanStruct(data, out)

	return
}

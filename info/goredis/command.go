package goredis

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func Select(num string) (err error) {
	_, err = conn().Do("SELECT", num)
	if err != nil {
		return
	}

	return
}

func Set(key string, value []byte) (err error) {
	_, err = conn().Do("SET", key, value)
	if err != nil {
		v := string(value)
		if len(v) > 15 {
			v = v[0:12] + "..."
		}
		return fmt.Errorf("error setting key %s to %s: %v", key, v, err)
	}
	return err
}

func Insert(key string, value interface{}) (err error) {
	_, err = conn().Do("SET", key, value)
	if err != nil {
		return fmt.Errorf("insert error %s: %v", key, err)
	}

	return
}

func Get(key string) (data []byte, err error) {
	data, err = redis.Bytes(conn().Do("GET", key))
	if err != nil {
		return data, fmt.Errorf("error getting key %s: %v", key, err)
	}

	return data, err
}

func GetUint(key string) (data uint64, err error) {
	data, err = redis.Uint64(conn().Do("GET", key))
	if err != nil {
		if err != nil {
			return data, fmt.Errorf("error getting key %s: %v", key, err)
		}
	}

	return data, err
}

func GetFloat64(key string) (data float64, err error) {
	data, err = redis.Float64(conn().Do("GET", key))
	if err != nil {
		if err != nil {
			return data, fmt.Errorf("error getting key %s: %v", key, err)
		}
	}

	return data, err
}

func Exists(key string) (bool, error) {
	exist, err := redis.Bool(conn().Do("EXISTS", key))
	if err != nil {
		return exist, fmt.Errorf("error checking if key %s exists: %v", key, err)
	}
	return exist, err
}

func Delete(key string) error {
	_, err := conn().Do("DEL", key)
	return err
}

func GetKeys(pattern string) ([]string, error) {
	iter := 0
	keys := []string{}
	for {
		arr, err := redis.Values(conn().Do("SCAN", iter, "MATCH", pattern))
		if err != nil {
			return keys, fmt.Errorf("error retrieving '%s' keys", pattern)
		}

		iter, _ = redis.Int(arr[0], nil)
		k, _ := redis.Strings(arr[1], nil)
		keys = append(keys, k...)

		if iter == 0 {
			break
		}
	}

	return keys, nil
}

func Increase(key string) (int, error) {
	return redis.Int(conn().Do("INCR", key))
}

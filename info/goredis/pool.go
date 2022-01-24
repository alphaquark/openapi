package goredis

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/gravision/alphaquark-upbit-api/logger"
	myutil "github.com/gravision/alphaquark-upbit-api/utils"
)

var (
	ErrNotExistEnvValue = errors.New("not exist env value")
	ErrFailedInitRedis  = errors.New("failed init redis")
)

var (
	pool *redis.Pool
)

const (
	DEFAULT_MAX_IDLE       = 3
	DEFAULT_TIMEOUT_SECOND = 180
)

func init() {
	// os.Setenv("REDIS_HOST", "127.0.0.1") // to Debug
	// os.Setenv("REDIS_PORT", "6379")      // to Debug
	// os.Setenv("REDIS_DB_NUMBER", "0")    // to Debug
	// os.Setenv("LOG_DEBUG_LEVEL", "INFO") // to Debug

	envHost := os.Getenv("REDIS_HOST")
	envPort := os.Getenv("REDIS_PORT")
	envDatabaseNumber := os.Getenv("REDIS_DB_NUMBER")

	if envHost == "" || envPort == "" || envDatabaseNumber == "" {
		logger.Panic("check redis environment")
	}

	redisAddr := envHost + ":" + envPort
	logger.Info(fmt.Sprintf("redis address: %s", redisAddr))

	pool = createPool(redisAddr)

	c := myutil.CleanupSignalHook()
	go func() {
		<-c
		pool.Close()
	}()
}

func createPool(server string) *redis.Pool {
	envMaxIdle := os.Getenv("REDIS_MAX_IDLE")

	var maxIdle int
	switch envMaxIdle {
	case "":
		maxIdle = DEFAULT_MAX_IDLE
		logger.Debug("Input environment value - Max Idle on Pool: Default, " + fmt.Sprintf("%d", DEFAULT_MAX_IDLE))
	default:
		maxIdle, _ = strconv.Atoi(envMaxIdle)
		logger.Debug(fmt.Sprintf("Input environment value - Max Idle on Pool: %d", maxIdle))
	}

	return &redis.Pool{
		MaxIdle:     maxIdle,
		IdleTimeout: DEFAULT_TIMEOUT_SECOND * time.Second,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", server)
			if err != nil {
				logger.Panic(fmt.Sprintf("failed to dial redis, %s", err.Error()))
				return nil, err
			}

			return conn, err
		},
	}
}

func conn() redis.Conn {
	return pool.Get()
}

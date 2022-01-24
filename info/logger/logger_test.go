package logger_test

import (
	"testing"

	logger "github.com/gravision/alphaquark-upbit-api/logger"
)

func TestLogger_Logging(t *testing.T) {
	// Logging
	msg := "this is test message."
	key1 := "key1"
	key2 := "key2"
	val1 := "value 1"
	val2 := "value 2"

	m := make(map[string]interface{})
	m[key1] = val1
	m[key2] = val2

	logger.Initialize("TRACE")

	// Trace level
	logger.Trace(msg)
	logger.TraceField(key1, val1, msg)
	logger.TraceFields(m, msg)

	// Debug level
	logger.Debug(msg)
	logger.DebugField(key1, val1, msg)
	logger.DebugFields(m, msg)

	// Info level
	logger.Info(msg)
	logger.InfoField(key1, val1, msg)
	logger.InfoFields(m, msg)

	// Warn level
	logger.Warn(msg)
	logger.WarnField(key1, val1, msg)
	logger.WarnFields(m, msg)

	// Error level
	logger.Error(msg)
	logger.ErrorField(key1, val1, msg)
	logger.ErrorFields(m, msg)

	// // Fatal level 은 exit 1 코드를 준다.
	// logger.Fatal(msg)
	// logger.FatalField(key1, val1, msg)
	// logger.FatalFields(m, msg)

	// // Panic level 은 panic error 를 준다.
	// logger.Panic(msg)
	// logger.PanicField(key1, val1, msg)
	// logger.PanicFields(m, msg)
}

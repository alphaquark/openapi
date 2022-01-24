package logger

import (
	"fmt"
	"os"

	logrus "github.com/sirupsen/logrus"
)

const (
	token = " | "
)

func init() {
	logLevel := os.Getenv("LOG_DEBUG_LEVEL")

	Initialize(logLevel)
}

func Initialize(log_level string) {
	logrus.SetFormatter(&logrus.TextFormatter{}) // Instead of ASCII Formatter
	logrus.SetOutput(os.Stdout)                  // Instead of `default` stderr

	switch log_level { // Set Logging level
	case "PANIC":
		// Highest level of severity. Logs and then calls panic with the message passed to Debug, Info, ...
		logrus.SetLevel(logrus.PanicLevel)
	case "FATAL":
		// Logs and then calls `logger.Exit(1)`. It will exit even if the logging level is set to Panic.
		logrus.SetLevel(logrus.FatalLevel)
	case "ERROR":
		// Used for errors that should definitely be noted.
		// Commonly used for hooks to send errors to an error tracking service.
		logrus.SetLevel(logrus.ErrorLevel)
	case "WARN":
		// Non-critical entries that deserve eyes.
		logrus.SetLevel(logrus.WarnLevel)
	case "INFO":
		// General operational entries about what's going on inside the application.
		logrus.SetLevel(logrus.InfoLevel)
	case "DEBUG":
		// Usually only enabled when debugging. Very verbose logging.
		logrus.SetLevel(logrus.DebugLevel)
	case "TRACE":
		// Designates finer-grained informational events than the Debug.
		logrus.SetLevel(logrus.TraceLevel)
	default:
		// Default loglevel > DebugLevel
		logrus.SetLevel(logrus.InfoLevel)
	}

	InfoFields(map[string]interface{}{
		"formatter": "JSON",
		"loglevel":  logrus.GetLevel(),
	}, "Initialized logger")
}

func Panic(msg interface{}) {
	message := fmt.Sprint(GetCaller(), token, msg)
	logrus.Panic(message)
}

func PanicField(key string, value interface{}, msg interface{}) {
	message := fmt.Sprint(GetCaller(), token, msg)
	logrus.WithField(key, value).Panic(message)
}

func PanicFields(datas map[string]interface{}, msg interface{}) {
	message := fmt.Sprint(GetCaller(), token, msg)
	fields := logrus.Fields{}

	for key, val := range datas {
		fields[key] = val
	}

	logrus.WithFields(fields).Panic(message)
}

func Fatal(msg interface{}) {
	message := fmt.Sprint(GetCaller(), token, msg)
	logrus.Fatal(message)
}

func FatalField(key string, value interface{}, msg interface{}) {
	message := fmt.Sprint(GetCaller(), token, msg)
	logrus.WithField(key, value).Fatal(message)
}

func FatalFields(datas map[string]interface{}, msg interface{}) {
	message := fmt.Sprint(GetCaller(), token, msg)
	fields := logrus.Fields{}

	for key, val := range datas {
		fields[key] = val
	}

	logrus.WithFields(fields).Fatal(message)
}

func Error(msg interface{}) {
	message := fmt.Sprint(GetCaller(), token, msg)
	logrus.Error(message)
}

func ErrorField(key string, value interface{}, msg interface{}) {
	message := fmt.Sprint(GetCaller(), token, msg)
	logrus.WithField(key, value).Error(message)
}

func ErrorFields(datas map[string]interface{}, msg interface{}) {
	message := fmt.Sprint(GetCaller(), token, msg)
	fields := logrus.Fields{}

	for key, val := range datas {
		fields[key] = val
	}

	logrus.WithFields(fields).Error(message)
}

func Warn(msg interface{}) {
	message := fmt.Sprint(GetCaller(), token, msg)
	logrus.Warn(message)
}

func WarnField(key string, value interface{}, msg interface{}) {
	message := fmt.Sprint(GetCaller(), token, msg)
	logrus.WithField(key, value).Warn(message)
}

func WarnFields(datas map[string]interface{}, msg interface{}) {
	message := fmt.Sprint(GetCaller(), token, msg)
	fields := logrus.Fields{}

	for key, val := range datas {
		fields[key] = val
	}

	logrus.WithFields(fields).Warn(message)
}

func Info(msg interface{}) {
	message := fmt.Sprint(GetCaller(), token, msg)
	logrus.Info(message)
}

func InfoField(key string, value interface{}, msg interface{}) {
	message := fmt.Sprint(GetCaller(), token, msg)
	logrus.WithField(key, value).Info(message)
}

func InfoFields(datas map[string]interface{}, msg interface{}) {
	message := fmt.Sprint(GetCaller(), token, msg)
	fields := logrus.Fields{}

	for key, val := range datas {
		fields[key] = val
	}

	logrus.WithFields(fields).Info(message)
}

func Debug(msg interface{}) {
	message := fmt.Sprint(GetCaller(), token, msg)
	logrus.Debug(message)
}

func DebugField(key string, value interface{}, msg interface{}) {
	message := fmt.Sprint(GetCaller(), token, msg)
	logrus.WithField(key, value).Debug(message)
}

func DebugFields(datas map[string]interface{}, msg interface{}) {
	message := fmt.Sprint(GetCaller(), token, msg)
	fields := logrus.Fields{}

	for key, val := range datas {
		fields[key] = val
	}

	logrus.WithFields(fields).Debug(message)
}

func Trace(msg interface{}) {
	message := fmt.Sprint(GetCaller(), token, msg)
	logrus.Info(message)
}

func TraceField(key string, value interface{}, msg interface{}) {
	message := fmt.Sprint(GetCaller(), token, msg)
	logrus.WithField(key, value).Trace(message)
}

func TraceFields(datas map[string]interface{}, msg interface{}) {
	message := fmt.Sprint(GetCaller(), token, msg)
	fields := logrus.Fields{}

	for key, val := range datas {
		fields[key] = val
	}

	logrus.WithFields(fields).Trace(message)
}

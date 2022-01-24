package logger

import (
	"runtime"
	"strconv"
)

func GetCaller() string {
	pc, _, _, _ := runtime.Caller(2)
	runtimeFunc := runtime.FuncForPC(pc)
	_, runtimeLine := runtimeFunc.FileLine(pc)
	runtimeName := runtimeFunc.Name()

	return runtimeName + ":" + strconv.Itoa(runtimeLine)
}

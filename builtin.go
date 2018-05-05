package goutils

import (
	"runtime"
)

// __FILE__, __func__, __LINE__
// fmt.Println(goutils.BuiltInfo())
func BuiltInfo() (file string, funcname string, funcline int){
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	return frame.File, frame.Function, frame.Line
}

package lib

import (
	"log"
	"runtime"
)

func PrintStackTrace(err error) {
	stackTrace := make([]byte, 1024)
	runtime.Stack(stackTrace, false)
	log.Printf("Stack trace:\n%s\n", stackTrace)
}

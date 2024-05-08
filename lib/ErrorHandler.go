package lib

import (
	"fmt"
	"runtime"
)

func PrintStackTrace(err error) {
	stackTrace := make([]byte, 1024)
	runtime.Stack(stackTrace, false)
	fmt.Printf("Stack trace:\n%s\n", stackTrace)
}

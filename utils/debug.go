package utils

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

// Debug func.
type Debug func(format string, v ...interface{})

var print = func(msg string) {
	fmt.Println(msg)
}

// InitDebug returns debug func.
func InitDebug() Debug {
	enable := os.Getenv("DEBUG") == "true"

	return func(format string, v ...interface{}) {
		if !enable {
			return
		}

		msg := fmt.Sprintf(format, v...)
		_, file, line, ok := runtime.Caller(1)
		if !ok {
			print(msg)
		} else {
			slash := strings.LastIndex(file, "forpay-sdk-go")

			fileInfo := fmt.Sprintf("%s:%d", file[slash+14:], line)
			print(fmt.Sprintf("%s: %s", fileInfo, msg))
		}
	}
}

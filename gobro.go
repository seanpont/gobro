package gobro

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

// If the error is not nil, exit with error code 1.
// Message is optional. Including more than one message will not have any result.
func CheckErr(err error, message ...string) {
	if err != nil {
		var msg string
		if len(message) > 0 {
			msg = message[0] + " "
		}
		errorMessage := caller() + msg + err.Error()
		fmt.Fprintf(os.Stderr, errorMessage+"\n")
		os.Exit(1)
	}
}

func LogErr(err error, message ...string) {
	if err != nil {
		var msg string
		if len(message) > 0 {
			msg = message[0] + " "
		}
		errorMessage := caller() + msg + err.Error()
		fmt.Fprintf(os.Stderr, errorMessage+"\n")
	}
}

func caller() string {
	var stack [4096]byte
	n := runtime.Stack(stack[:], false)
	caller := strings.Split(string(stack[:n]), "\n")[6]
	caller = strings.Trim(caller, " \t")
	return strings.Split(caller, " ")[0] + ": "
}

// ===== PRIMITIVE UTILS =====================================================

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

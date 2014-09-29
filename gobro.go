package gobro

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

// If the error is not nil, exit with error code 1.
// Message is optional. Including more than one message will not have any result.
func ExitOnError(err error, message ...string) {
	if err != nil {
		errorMessage := caller() + err.Error()
		if len(message) > 0 {
			errorMessage += message[0]
		}
		fmt.Fprintf(os.Stderr, errorMessage+"\n")
		os.Exit(1)
	}
}

func caller() string {
	var stack [4096]byte
	n := runtime.Stack(stack[:], false)
	caller := strings.Split(string(stack[:n]), "\n")[6]
	caller = strings.Trim(caller, " \t")
	return strings.Split(caller, " ")[0] + ": "
}

package gobro

import (
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strings"
)

// If the error is not nil, exit with error code 1.
// Message is optional. Including more than one message will not have any result.
func ExitOnError(err error, message ...string) {
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

func LogError(err error, message ...string) {
	if err != nil {
		var msg string
		if len(message) > 0 {
			msg = message[0] + " "
		}
		errorMessage := caller() + msg + err.Error()
		fmt.Fprintf(os.Stderr, errorMessage+"\n")
	}
}

func TrimAll(items []string) {
	for i, item := range items {
		items[i] = strings.Trim(item, " \n\r\t")
	}
}

func caller() string {
	var stack [4096]byte
	n := runtime.Stack(stack[:], false)
	caller := strings.Split(string(stack[:n]), "\n")[6]
	caller = strings.Trim(caller, " \t")
	return strings.Split(caller, " ")[0] + ": "
}

// ===== COMMAND MAPPER ======================================================

type CommandMap struct {
	// CommandMap holds a map of names to functions. Useful for handling
	// control flow in main functions writing a ton of if this else that or
	// using flag, which I find sub-optimal
	commandMap map[string]func([]string)
}

func NewCommandMap(functions ...func(args []string)) (commandMap CommandMap) {
	// Returns a new CommandMap with the functions mapped to their names.
	commandMap.commandMap = make(map[string]func([]string))

	for _, fn := range functions {
		name := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
		name = strings.Split(name, ".")[1] // main.command becomes command
		commandMap.commandMap[name] = fn
	}

	return
}

func (cm CommandMap) Run(args []string) {

	if len(os.Args) == 1 {
		fmt.Printf("Usage: %s [options] <command> [<args>]\n\n", args[0])
		fmt.Println("Available commands:")
		for k, _ := range cm.commandMap {
			fmt.Printf("   %s\n", k)
		}
		os.Exit(1)
	}

	command := os.Args[1]
	fn := cm.commandMap[command]
	if fn == nil {
		fmt.Fprintf(os.Stderr, "Command '%s' not found\n", command)
		os.Exit(1)
	}

	fn(os.Args[2:])
}

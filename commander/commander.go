// Command line utility functions
package commander

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"runtime"
	"sort"
	"strings"
)

// ===== COMMAND MAPPER ======================================================

type FuncDesc struct {
	Fn   func([]string)
	Desc string
}

type CommandMap struct {
	// CommandMap holds a map of names to functions. Useful for handling
	// control flow in main functions writing a ton of if this else that or
	// using flag, which I find sub-optimal
	funcMap map[string]FuncDesc
}

func NewCommandMap(functions ...func(args []string)) *CommandMap {
	// Returns a new CommandMap with the functions mapped to their names.
	// Usage: gobro.NewCommandMap(configure, doSomething).Run(os.Args)
	commandMap := new(CommandMap)
	commandMap.funcMap = make(map[string]FuncDesc)

	for _, fn := range functions {
		name := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
		nameStart := strings.LastIndex(name, ".") + 1
		// If the method is part of a struct, the function name will be like this:
		// github.com/seanpont/gobro/commander.TestFunctionProvider.(github.com/seanpont/gobro/commander.foo)·fm
		nameEnd := strings.LastIndex(name, ")")
		if nameEnd == -1 {
			nameEnd = len(name)
		}
		cleanName := name[nameStart:nameEnd]
		commandMap.funcMap[cleanName] = FuncDesc{Fn: fn}
	}

	return commandMap
}

func (cm *CommandMap) Add(name string, fn func([]string), desc ...string) {
	if len(desc) > 0 {
		cm.funcMap[name] = FuncDesc{Fn: fn, Desc: desc[0]}
	} else {
		cm.funcMap[name] = FuncDesc{Fn: fn}
	}
}

func (cm *CommandMap) Commands() []string {
	commands := make([]string, 0, len(cm.funcMap))
	for k, _ := range cm.funcMap {
		commands = append(commands, k)
	}
	sort.Strings(commands)
	return commands
}

func (cm *CommandMap) Run(args []string) {
	// Run the function corresponding to the first argument in args
	// You're probably going to want to pass in os.Args
	cmd := ""
	options := make([]string, 0)
	if len(args) > 1 {
		cmd = args[1]
		options = args[2:]
	}

	fn := cm.funcMap[cmd]
	if fn.Fn != nil {
		fn.Fn(options)
	} else {
		fmt.Printf("Usage: %s [options] <command> [<args>]\n\n", args[0])
		fmt.Println("Available commands:")
		for _, k := range cm.Commands() {
			v := cm.funcMap[k]
			fmt.Printf("   %-10s  %-10s\n", k, v.Desc)
		}
		fmt.Println("")
	}
}

func CheckArgs(args []string, numArgs int, message string, a ...interface{}) {
	// Helper function for verifying that the args are correct
	if len(args) != numArgs {
		fmt.Fprintf(os.Stderr, message+"\n", a...)
		os.Exit(1)
	}
}

// ===== COMMAND LINE TOOLS ==================================================

func Prompt(query string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(query)
	line, _, err := reader.ReadLine()
	if err != nil {
		return "", err
	}
	return string(line), nil
}

package commander

import (
	"fmt"
	"github.com/seanpont/assert"
	"testing"
)

type TestFunctionProvider struct {
	method string
	args   []string
}

func (tfp *TestFunctionProvider) foo(args []string) {
	fmt.Printf("Foo called on %v", tfp)
	tfp.method = "foo"
	tfp.args = args
}

func (tfp *TestFunctionProvider) bar(args []string) {
	tfp.method = "bar"
	tfp.args = args
}

func fizzBuzz(args []string) {}

func TestCommander(t *testing.T) {
	assert := assert.Assert(t)
	tfp := new(TestFunctionProvider)
	commandMap := NewCommandMap(fizzBuzz, tfp.foo, tfp.bar)
	commandMap.Add("fuzz", fizzBuzz, "FizzBuzz description")
	assert.Equal(len(commandMap.funcMap), 4)
	assert.Same(commandMap.funcMap["fizzBuzz"].Fn, fizzBuzz)
	assert.Same(commandMap.funcMap["foo"].Fn, tfp.foo)
	assert.Equal(commandMap.funcMap["foo"].Desc, "")
	assert.Same(commandMap.funcMap["fuzz"].Fn, fizzBuzz)
	assert.Equal(commandMap.funcMap["fuzz"].Desc, "FizzBuzz description")

	// The first argument is always the name of the program
	args := []string{"commander_test", "foo", "arg1", "arg2"}
	commandMap.Run(args)
	assert.Equal(tfp.method, "foo")
	assert.Equal(tfp.args, args[2:])
}

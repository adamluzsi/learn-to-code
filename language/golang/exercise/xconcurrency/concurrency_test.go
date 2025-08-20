package xconcurrency_test

import (
	"context"
	"go/ast"
	exercise "golang-exercise"
	"golang-exercise/xconcurrency"
	"testing"
	"time"

	"go.llib.dev/testcase/assert"
	"go.llib.dev/testcase/pp"
)

const chanExplanation = `
Let's learn about channels!

A channel in Go is like a tube
that requires both sender and receiver to be ready to exchange data.
If someone sends an object into the tube (like putting a message in),
the other person must wait in the tube until they can take it out of their hands.
This ensures tasks happen in order — if one isn't ready, the other has to wait until they are.

In this exercise, Let's use the <-chan and chan<-
take out the value from the input and place it back into the output channel`

func TestExercise1(t *testing.T) {
	exercise.Description(t, chanExplanation)

	var (
		input  = make(chan int)
		output = make(chan int)
		exp    = exercise.Random().Int()
		got    int
	)

	go xconcurrency.HowToWorkWithChannels(input, output)

	assert.Within(t, time.Second/2, func(ctx context.Context) {
		input <- exp
	}, "in order to receive values from a channel,",
		"you can use either the `for val := range channelvarname {}` keywords",
		"or `variable := <-channelvarname`")

	assert.Within(t, time.Second/2, func(ctx context.Context) {
		got = <-output
	}, "in order to send a value through a channel,",
		"you can use `channelname <- value`")

	assert.Equal(t, exp, got, "You are expected to send back on the output channel whatever you got from the input channel.")
}

const descGoroutine = `
Go's goroutines are an efficient way to run code concurrently, much like multitasking in everyday life.
You don't always do tasks in parallel, but you can work on them seemingly at the same time.
For example, making a phone call while using your hands for another task.
You switch between these activities, allowing them to progress together.
However, the more tasks you handle, the less time you have on each before switching to the next.

Modern computers often have multiple CPUs, which can sometimes allow work to be done in parallel.
But this isn't guaranteed, so it's best not to rely on it in your design.
Working with concurrency can offer benefits if your code happens to run tasks simultaneously.

Let's try creating a goroutine together! Use the "go" keyword.

This example for example executes an anonymous function on a seperate goroutine:

` + "```go" + `
go func() {
	/* function that will run as a new goroutine */
}()
` + "```" + `

And in this example we execute a named function on a new goroutine:

` + "```go" + `
function MyFunc() {}

func main() {
	go MyFunc()
}
` + "```" + `
`

func TestExercise2(t *testing.T) {
	exercise.Description(t, descGoroutine)

	fn := exercise.GetFunc(t, "concurrency.go", "MakeAGoroutine")

	pp.PP(fn)
	ast.Inspect(fn, func(n ast.Node) bool {
		return true
	})
}

const descChannelsAreBlocking = `
To understand more how channels behave, let's take the following example.

Given you have two channels:
- one where need to put a value inside
- one where you need to take out a value

Since a channel is like having your hand holding out an object ,and waiting for someone to take it out from your hand
it will block the current programflow's execution
to avoid that, you could use a goroutine.
`

func TestExercise3(t *testing.T) {
	exercise.Description(t, descChannelsAreBlocking)

	var (
		input  = make(chan int)
		output = make(chan int)
		exp    = exercise.Random().Int()
		got    int
	)

	assert.Within(t, time.Second, func(ctx context.Context) {
		xconcurrency.ChannelsAreBlocking(input, output)
		input <- exp
		got = <-output
	})

	assert.Equal(t, exp, got,
		"You are expected to send back on the output channel whatever you got from the input channel.")
}

package main

import (
	"flag"
	"fmt"
	"os"

	todo "github.com/FilledEther20/c.l.i.g.o"
)

const (
	todoFile = ".todos.json"
)

// function for error handling
func handleError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
func main() {
	add := flag.Bool("add", false, "add a new list")

	complete := flag.Int("complete", 0, "mark a task as completed through index")
	flag.Parse()

	todos := &todo.Todos{}

	if err := todos.Load(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	//the flow would be
	//1) perform the operation
	//2) store the state of the list
	switch {
	case *add:
		todos.Add("Test todo")
		err := todos.Store(todoFile)
		handleError(err)
	case *complete > 0:
		err := todos.Complete(*complete)
		handleError(err)
		err := todos.Store(todoFile)
		handleError(err)
	default:
		fmt.Fprintln(os.Stdout, "invalid command")
		os.Exit(1)
	}
}

package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

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

	delete := flag.Int("delete", 0, "delete a task through it's id")

	list := flag.Bool("list", false, "prints the entire todo members")
	flag.Parse()

	todos := &todo.Todos{}

	if err := todos.Load(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	//The flow is:
	//1) perform the operation
	//2) store the state of the list
	switch {
	case *add:
		input, err := getInput(os.Stdin, flag.Args()...)
		handleError(err)

		todos.Add(input)

		err = todos.Store(todoFile)
		handleError(err)

	case *complete > 0:
		err := todos.Complete(*complete)
		handleError(err)
		err = todos.Store(todoFile)
		handleError(err)

	case *delete > 0:
		err := todos.Deletion(*delete)
		handleError(err)
		err = todos.Store(todoFile)
		handleError(err)

	case *list:
		todos.Print()
	default:
		fmt.Fprintln(os.Stdout, "invalid command")
		os.Exit(1)
	}
}

//function to take input for the add functionality

func getInput(r io.Reader, args ...string) (string, error) {
	//The args and r provide two separate modes for the input to be accessed through
	//if input at io.Reader is not given then the input can be taken directly through args
	//else if input is provided then no need for arguments
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	scanner := bufio.NewScanner(r)
	scanner.Scan()

	if err := scanner.Err(); err != nil {
		return "", err
	}

	text := scanner.Text()
	if len(text) == 0 {
		return "", errors.New("task cannot be empty")
	}
	return text, nil
}

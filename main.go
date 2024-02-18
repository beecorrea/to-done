package main

import (
	"fmt"

	"github.com/beecorrea/to-done/pkg/todo"
)

var files = []string{"./pkg/todo/reader_test.go", "./data/test.go"}

func main() {
	errs := make(map[string]error, 0)
	todos := make([]string, 0)

	for _, target := range files {
		t, err := todo.TodosFrom(target)
		if err != nil {
			errs[target] = err
		} else {
			todos = append(todos, t...)
		}
	}

	if len(errs) > 0 {
		for f, err := range errs {
			fmt.Printf("couldn't process %s: %v", f, err)
		}
	}

	fmt.Printf("To-do:\n\n")
	for _, t := range todos {
		fmt.Printf("\t- %v\n", t)
	}
}

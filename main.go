package main

import (
	"fmt"

	"github.com/beecorrea/to-done/pkg/reader"
	"github.com/beecorrea/to-done/pkg/todo"
)

var dir = "."

func main() {
	errs := make(map[string]error, 0)
	todos := make([]string, 0)
	files, err := reader.GetFilesInDir(dir)
	if err != nil {
		panic(err)
	}

	for _, target := range files {
		t, err := todo.TodosFrom(target)
		if err != nil && err != todo.ErrNoTodos {
			errs[target] = err
		} else {
			todos = append(todos, t...)
		}
	}

	if len(errs) > 0 {
		for f, err := range errs {
			fmt.Printf("couldn't process %s: %v\n", f, err)
		}
	}

	fmt.Printf("%d to-dos:\n\n", len(todos))
	for _, t := range todos {
		fmt.Printf("\t- %v\n", t)
	}
}

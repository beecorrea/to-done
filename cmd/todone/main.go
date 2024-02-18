package main

import (
	"fmt"

	"github.com/beecorrea/to-done/internal/controllers"
	"github.com/beecorrea/to-done/pkg/reader"
	"github.com/beecorrea/to-done/pkg/todo"
)

var dir = "."

func main() {
	errs := make(map[string]error, 0)
	todos := make(map[string][]reader.LineInfo, 0)
	files, err := reader.GetFilesInDir(dir)
	if err != nil {
		panic(err)
	}

	numTodos := 0
	for _, target := range files {
		t, err := controllers.TodosFrom(target)
		if err != nil && err != todo.ErrNoTodos {
			errs[target] = err
		} else if len(t) > 0 {
			todos[target] = t
			numTodos += len(t)
		}
	}

	if len(errs) > 0 {
		for f, err := range errs {
			fmt.Printf("couldn't process %s: %v\n", f, err)
		}
	}

	fmt.Printf("%d to-dos:\n\n", numTodos)
	for filename, t := range todos {
		fmt.Printf(" [%s]\n", filename)
		for _, line := range t {
			fmt.Printf("   (line %d): %s\n", line.Number, line.Content)
		}
		fmt.Printf("\n")
	}
}

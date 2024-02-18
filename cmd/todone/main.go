package main

import (
	"fmt"

	"github.com/beecorrea/to-done/internal/controllers"
)

var dir = "."

func main() {
	todos, errs := controllers.Todos(dir)
	if len(errs) > 0 {
		for f, err := range errs {
			fmt.Printf("couldn't process %s: %v\n", f, err)
		}
	}

	numTodos := 0
	for _, v := range todos {
		numTodos += len(v)
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

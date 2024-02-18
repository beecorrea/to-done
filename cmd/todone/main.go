package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/beecorrea/to-done/internal/controllers"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

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
		prettyFile := strings.Replace(filename, dir+"/", "", -1)
		fmt.Printf(" [%s]\n", prettyFile)
		for _, line := range t {
			fmt.Printf("   (line %d): %s\n", line.Number, line.Content)
		}
		fmt.Printf("\n")
	}
}

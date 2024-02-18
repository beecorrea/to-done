package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/beecorrea/to-done/internal/controllers"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "todone [folder]",
	Args:  cobra.MaximumNArgs(1),
	Short: "View all todos in your project's source files.",
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {
	folder, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	if len(args) > 0 {
		folder = args[0]
	}

	todos, errs := controllers.Todos(folder)
	if len(errs) > 0 {
		for f, err := range errs {
			fmt.Printf("couldn't process %s: %v\n", f, err)
		}
	}

	numTodos := 0
	for _, v := range todos {
		numTodos += len(v)
	}

	for filename, t := range todos {
		prettyFile := strings.Replace(filename, folder+"/", "", -1)
		fmt.Printf(" [%s]\n", prettyFile)
		for _, line := range t {
			fmt.Printf("   (line %d): %s\n", line.Number, line.Content)
		}
		fmt.Printf("\n")
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

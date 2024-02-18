package todo

import (
	"fmt"
	"regexp"
)

// todo: find a way to make beginsWith regex work.
// todo: allow passing single line comments as a parameter.
var opTodo = `// (?i)todo:(.+)`

// Todo: convert to custom error
var ErrNoTodos = fmt.Errorf("file has no todos")

// Tries to build a regex using OpTodo
func OpTodo() (*regexp.Regexp, error) {
	return regexp.Compile(opTodo)
}

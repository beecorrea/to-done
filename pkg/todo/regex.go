package todo

import (
	"fmt"
	"regexp"
)

// todo: allow passing single line comments as a parameter.
var opTodo = `(?mi)^\s*//\s*todo:\s*(.+)`
var ROpTodo = regexp.MustCompile(opTodo)

// Todo: convert to custom error
var ErrNoTodos = fmt.Errorf("file has no todos")

// Tries to build a regex using OpTodo
func OpTodo() (*regexp.Regexp, error) {
	return regexp.Compile(opTodo)
}

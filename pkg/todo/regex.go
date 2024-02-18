package todo

import (
	"fmt"
	"regexp"
)

var opTodo = `// (?i)todo:(.+)`
var ErrNoTodos = fmt.Errorf("file has no todos")

// Tries to build a regex using OpTodo
func OpTodo() (*regexp.Regexp, error) {
	return regexp.Compile(opTodo)
}

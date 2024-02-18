package todo

import (
	"fmt"
	"regexp"
	"strings"
)

var opTodo = `// (?i)todo:(.+)`
var ErrNoTodos = fmt.Errorf("file has no todos")

func TodosFrom(path string) ([]string, error) {
	r, err := OpTodo()
	if err != nil {
		return nil, err
	}

	ct, err := GetFileContents(path)
	if err != nil {
		return nil, err
	}

	todos := FindTodos(r, ct)
	if len(todos) < 1 {
		return nil, ErrNoTodos
	}
	return todos, nil
}

// Tries to build a regex using OpTodo
func OpTodo() (*regexp.Regexp, error) {
	return regexp.Compile(opTodo)
}

// Matches the first todo using a regex.
func FindTodo(r *regexp.Regexp, s string) string {
	matches := r.FindStringSubmatch(s)
	first := matches[1]
	return strings.Trim(first, " ")
}

// Matches all todos in a string using the OpTodo regex.
func FindTodos(r *regexp.Regexp, s string) []string {
	groups := r.FindAllStringSubmatch(s, -1)
	trimmed := make([]string, 0)

	for _, g := range groups {
		for _, m := range g[1:] {
			trim := strings.Trim(m, " ")
			trimmed = append(trimmed, trim)
		}
	}

	return trimmed
}

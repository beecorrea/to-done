package todo

import (
	"regexp"
	"strings"
)

// Matches the first todo using a regex.
func FindTodo(r *regexp.Regexp, s string) string {
	matches := r.FindStringSubmatch(s)
	first := matches[1]
	return strings.Trim(first, " ")
}

// Matches all todos in a string using the OpTodo regex.
func FindTodos(r *regexp.Regexp, s string) []string {
	groups := r.FindAllStringSubmatch(s, -1)
	if len(groups) == 0 {
		return nil
	}

	trimmed := make([]string, 0)
	// todo: should be refactored?
	for _, g := range groups {
		for _, m := range g[1:] {
			trim := strings.Trim(m, " ")
			trimmed = append(trimmed, trim)
		}
	}

	return trimmed
}

package controllers

import (
	"github.com/beecorrea/to-done/pkg/reader"
	"github.com/beecorrea/to-done/pkg/todo"
)

func TodosFrom(path string) ([]reader.LineInfo, error) {
	lines, err := reader.WalkFile(path, findTodoWalk)
	if err != nil {
		return nil, err
	}
	if len(lines) == 0 {
		return nil, todo.ErrNoTodos
	}

	return lines, nil
}

func findTodoWalk(s string) string {
	return todo.FindTodo(todo.ROpTodo, s)
}

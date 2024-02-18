package controllers

import (
	"github.com/beecorrea/to-done/pkg/reader"
	"github.com/beecorrea/to-done/pkg/todo"
)

func todosFrom(path string) ([]reader.LineInfo, error) {
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

func Todos(root string) (map[string][]reader.LineInfo, map[string]error) {
	errs := make(map[string]error, 0)
	todos := make(map[string][]reader.LineInfo, 0)
	files, err := reader.GetFilesInDir(root)
	if err != nil {
		panic(err)
	}

	numTodos := 0
	for _, target := range files {
		t, err := todosFrom(target)
		if err != nil && err != todo.ErrNoTodos {
			errs[target] = err
		} else if len(t) > 0 {
			todos[target] = t
			numTodos += len(t)
		}
	}

	return todos, errs
}

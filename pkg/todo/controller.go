package todo

import "github.com/beecorrea/to-done/pkg/reader"

func TodosFrom(path string) ([]string, error) {
	r, err := OpTodo()
	if err != nil {
		return nil, err
	}

	ct, err := reader.GetFileContents(path)
	if err != nil {
		return nil, err
	}

	todos := FindTodos(r, ct)
	// todo: change to len(todos) == 0
	if len(todos) < 1 {
		return nil, ErrNoTodos
	}
	return todos, nil
}

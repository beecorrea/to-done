package todo

import (
	"fmt"
	"strings"
	"testing"
)

func TestOpTodoCompiles(t *testing.T) {
	if _, err := OpTodo(); err != nil {
		t.Fatalf("should compile OpTodo, but got error %s", err.Error())
	}
}

func TestMatchTodoRegex(t *testing.T) {
	text := "Should pass."
	todo := fmt.Sprintf("// TODO: %s", text)
	r, _ := OpTodo()

	match := FindTodo(r, todo)
	if match != text {
		t.Fatalf("should match and return '%s', got '%s'", text, match)
	}
}

func TestReadManyTodos(t *testing.T) {
	texts := []string{"This is a todo.", "Another todo."}
	todos := strings.Join([]string{useTodo(texts[0]), useTodo(texts[1])}, "\n")
	r, _ := OpTodo()

	matches := FindTodos(r, todos)
	for i := range matches {
		if matches[i] != texts[i] {
			t.Fatalf("todo should match '%s', got '%s' instead", texts[i], matches[i])
		}
	}
}

func useTodo(todo string) string {
	return fmt.Sprintf("// TODO: %s", todo)
}

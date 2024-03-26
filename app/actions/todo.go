package actions

import (
	"github.com/jameserrington/neon-arbor/app/models"
	"github.com/jameserrington/neon-arbor/app/storage"
)

func CreateTodo(task string) (models.Todo, error) {
	db := storage.GetDB()
	query := "INSERT INTO todos (task, completed) VALUES ($1, $2) RETURNING id, task, completed, created_at, updated_at"

	var todo models.Todo
	err := db.QueryRow(query, task, false).Scan(&todo.Id, &todo.Task, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		return todo, err
	}

	return todo, nil
}

func GetAllTodos() ([]models.Todo, error) {
	db := storage.GetDB()
	query := "SELECT id, task, completed, created_at, updated_at FROM todos ORDER BY created_at ASC"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	todos := make([]models.Todo, 0, 10)
	for rows.Next() {
		var todo models.Todo
		err := rows.Scan(&todo.Id, &todo.Task, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt)
		if err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	return todos, nil
}

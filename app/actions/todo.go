package actions

import (
	"github.com/jameserrington/neon-arbor/app/models"
	"github.com/jameserrington/neon-arbor/app/storage"
)

func CreateTodo(task string) (models.Todo, error) {
	db := storage.GetDB()
	query := "INSERT INTO todos (task, completed) VALUES ($1, $2) RETURNING *"

	var todo models.Todo
	err := db.QueryRow(query, task, false).Scan(&todo.Id, &todo.Task, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		return todo, err
	}

	return todo, nil
}

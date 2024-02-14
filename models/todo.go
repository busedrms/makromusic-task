package models

import (
	"context"
	"makromusic-task/utils"
	"strconv"

	"github.com/jackc/pgx/v4"
)

type Todo struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func GetTodoList(pageNum, pageSize int) ([]Todo, error) {
	var todos []Todo
	offset := (pageNum - 1) * pageSize
	query := "SELECT id, user_id, description, completed FROM todos ORDER BY id LIMIT $1 OFFSET $2"
	rows, err := utils.DB.Query(context.Background(), query, pageSize, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.ID, &todo.UserID, &todo.Description, &todo.Completed)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}
func (t *Todo) Create() error {
	query := "INSERT INTO todos (user_id, description, completed) VALUES ($1, $2, $3) RETURNING id"
	err := utils.DB.QueryRow(context.Background(), query, t.UserID, t.Description, t.Completed).Scan(&t.ID)
	if err != nil {
		return err
	}

	return nil
}
func MarkTodoAsComplete(todoID string) error {
	id, err := strconv.Atoi(todoID)
	if err != nil {
		return err
	}

	query := "UPDATE todos SET completed = true WHERE id = $1"
	_, err = utils.DB.Exec(context.Background(), query, id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return err
		}
		return err
	}

	return nil
}

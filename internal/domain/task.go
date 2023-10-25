package domain

import (
	"todo/models"
)

type Task struct {
	TaskID      int             `json:"task_id"`
	UserID      int             `json:"user_id"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Priority    models.Priority `json:"priority"`
	DueDate     DueDate         `json:"due_date"`
	Completed   bool            `json:"completed"`
}

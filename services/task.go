package services

import (
	"time"
	"todo/models"
)

type Task struct {
	TaskID      int
	UserID      int
	Title       string
	Description string
	Priority    models.Priority
	DueDate     time.Time
	Completed   bool
}

package services

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"log"
	"reflect"
	"strings"
	"time"
	"todo/models"
)

// type DueDate time.Time
type DueDate struct {
	time.Time
}

type Task struct {
	TaskID      int             `json:"task_id"`
	UserID      int             `json:"user_id"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Priority    models.Priority `json:"priority"`
	DueDate     DueDate         `json:"due_date"`
	Completed   bool            `json:"completed"`
}

func (d *DueDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	//*d = DueDate(t)
	d.Time = t
	return nil
}

func (d DueDate) MarshalJSON() ([]byte, error) {
	//return json.Marshal(time.Time(d))
	return json.Marshal(d.Time)
}

func (d *DueDate) Scan(value interface{}) error {
	if v, ok := value.(time.Time); ok {
		d.Time = v
		//*d = DueDate(v)
		return nil
	}
	return errors.New("Invalid DueDate value")
}

func (d DueDate) Value() (driver.Value, error) {
	if d.IsZero() {
		return nil, nil
	}
	return d.Time, nil
}

func (t *Task) GetAllTasks() ([]*Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT * FROM tasks`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var tasks []*Task
	for rows.Next() {
		var task Task

		s := reflect.ValueOf(&task).Elem()
		numCols := s.NumField()
		columns := make([]interface{}, numCols)
		for i := 0; i < numCols; i++ {
			field := s.Field(i)
			columns[i] = field.Addr().Interface()
		}

		err = rows.Scan(columns...)
		if err != nil {
			log.Fatalf("error while scanning rows (GetAllTasks):", err)
		}
		tasks = append(tasks, &task)
	}
	return tasks, nil
}

func (t *Task) GetTaskByID(taskID int) (*Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT * FROM tasks WHERE task_id = $1`

	var task Task

	row := db.QueryRowContext(ctx, query, taskID)
	err := row.Scan(
		&task.UserID,
		&task.TaskID,
		&task.Title,
		&task.Description,
		&task.Priority,
		&task.DueDate,
		&task.Completed,
	)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (t *Task) GetTasksByPriority(priority models.Priority) ([]*Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT * FROM tasks WHERE priority = $1`

	rows, err := db.QueryContext(ctx, query, priority)
	if err != nil {
		return nil, err
	}

	var tasks []*Task
	for rows.Next() {
		var task Task

		s := reflect.ValueOf(&task).Elem()
		numCols := s.NumField()
		columns := make([]interface{}, numCols)
		for i := 0; i < numCols; i++ {
			field := s.Field(i)
			columns[i] = field.Addr().Interface()
		}

		err = rows.Scan(columns...)
		if err != nil {
			log.Fatalf("error while scanning rows (GetTasksByPriority):", err)
		}
		tasks = append(tasks, &task)
	}
	return tasks, nil
}

func (t *Task) CreateTask(task Task) (*Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `INSERT INTO tasks (user_id, title, description, priority, due_date, completed) 
			  VALUES ($1, $2, $3, $4, $5, $6) RETURNING *`

	_, err := db.ExecContext(
		ctx,
		query,
		task.UserID,
		task.Title,
		task.Description,
		task.Priority,
		task.DueDate,
		task.Completed,
	)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (t *Task) UpdateTask(taskID int, task Task) (*Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
			UPDATE tasks
			SET
			    title = $1,
			    description = $2,
			    priority = $3,
			    due_date = $4,
			    completed = $5
			WHERE task_id = $6
     		`

	_, err := db.ExecContext(
		ctx,
		query,
		task.Title,
		task.Description,
		task.Priority,
		task.DueDate,
		task.Completed,
		taskID,
	)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (t *Task) DeleteTask(taskID int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `DELETE FROM tasks WHERE task_id = $1`
	_, err := db.ExecContext(ctx, query, taskID)
	if err != nil {
		return err
	}
	return nil
}

package repository

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"todo/internal/domain"
	"todo/models"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return Repository{db: db}
}

func (r Repository) GetAllTasks(ctx context.Context) ([]domain.Task, error) {
	rows, err := r.db.QueryContext(
		ctx,
		`
			select *
		    from tasks 
		`,
	)
	if err != nil {
		return nil, fmt.Errorf("getting all tasks: %w", err)
	}

	var tasks []domain.Task
	for rows.Next() {
		var task domain.Task

		s := reflect.ValueOf(&task).Elem()
		numCols := s.NumField()
		columns := make([]interface{}, numCols)
		for i := 0; i < numCols; i++ {
			field := s.Field(i)
			columns[i] = field.Addr().Interface()
		}

		err = rows.Scan(columns...)
		if err != nil {
			return nil, fmt.Errorf("error while scanning rows: %w", err)
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r Repository) GetTask(ctx context.Context, id int) (domain.Task, error) {
	var task domain.Task

	err := r.db.QueryRowContext(
		ctx,
		`
			select *
			from tasks
			where task_id = $1
		`,
		id,
	).Scan(
		&task.UserID,
		&task.TaskID,
		&task.Title,
		&task.Description,
		&task.Priority,
		&task.DueDate,
		&task.Completed,
	)

	if err != nil {
		return domain.Task{}, fmt.Errorf("getting task by id=%d: %w", id, err)
	}

	return task, nil
}

func (r Repository) GetTasksByPriority(ctx context.Context, priority models.Priority) ([]domain.Task, error) {
	rows, err := r.db.QueryContext(
		ctx,
		`
			select *
			from tasks
			where priority = $1
		`,
		priority,
	)
	if err != nil {
		return nil, fmt.Errorf("getting tasks by priority: %w", err)
	}

	var tasks []domain.Task
	for rows.Next() {
		var task domain.Task

		s := reflect.ValueOf(&task).Elem()
		numCols := s.NumField()
		columns := make([]interface{}, numCols)
		for i := 0; i < numCols; i++ {
			field := s.Field(i)
			columns[i] = field.Addr().Interface()
		}

		err = rows.Scan(columns...)
		if err != nil {
			return nil, fmt.Errorf("error while scanning rows: %w", err)
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r Repository) CreateTask(ctx context.Context, task domain.Task) (int, error) {
	var id int
	err := r.db.QueryRowContext(
		ctx,
		`
			insert into tasks (user_id, title, description, priority, due_date, completed) 
			values ($1, $2, $3, $4, $5, $6) returning task_id
		`,
		task.UserID,
		task.Title,
		task.Description,
		task.Priority,
		task.DueDate,
		task.Completed,
	).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("creating task: %w", err)
	}

	return id, nil
}

func (r Repository) UpdateTask(ctx context.Context, task domain.Task) (int, error) {
	var id int
	err := r.db.QueryRowContext(
		ctx,
		`
			update tasks
			set
			    title = $1,
			    description = $2,
			    priority = $3,
			    due_date = $4,
			    completed = $5
			where task_id = $6
		`,
		task.Title,
		task.Description,
		task.Priority,
		task.DueDate,
		task.Completed,
		task.TaskID,
	).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("updating task with id=%d: %w", task.TaskID, err)
	}

	return id, nil
}

func (r Repository) DeleteTask(ctx context.Context, id int) error {
	_, err := r.db.ExecContext(
		ctx,
		`
			delete from tasks
			where task_id = $1
		`,
		id,
	)
	if err != nil {
		return fmt.Errorf("deleting task with id=%d: %w", err)
	}
	return nil
}

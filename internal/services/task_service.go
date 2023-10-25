package services

import (
	"context"
	"fmt"
	"todo/internal/domain"
	"todo/models"
)

type Repository interface {
	GetAllTasks(ctx context.Context) ([]domain.Task, error)
	GetTask(ctx context.Context, id int) (domain.Task, error)
	GetTasksByPriority(ctx context.Context, priority models.Priority) ([]domain.Task, error)
	CreateTask(ctx context.Context, task domain.Task) (int, error)
	UpdateTask(ctx context.Context, task domain.Task) (int, error)
	DeleteTask(ctx context.Context, id int) error
}

type TaskService struct {
	repo Repository
}

func NewTaskService(repo Repository) TaskService {
	return TaskService{repo: repo}
}

func (s TaskService) GetTasks(ctx context.Context) ([]domain.Task, error) {
	tasks, err := s.repo.GetAllTasks(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting all tasks: %w", err)
	}

	return tasks, nil
}

func (s TaskService) GetTask(ctx context.Context, id int) (domain.Task, error) {
	task, err := s.repo.GetTask(ctx, id)
	if err != nil {
		return domain.Task{}, fmt.Errorf("getting task: %w", err)
	}

	return task, nil
}

func (s TaskService) GetTasksByPriority(ctx context.Context, priority models.Priority) ([]domain.Task, error) {
	tasks, err := s.repo.GetTasksByPriority(ctx, priority)
	if err != nil {
		return nil, fmt.Errorf("getting tasks: %w", err)
	}

	return tasks, nil
}

func (s TaskService) CreateTask(ctx context.Context, task domain.Task) (int, error) {
	id, err := s.repo.CreateTask(ctx, task)
	if err != nil {
		return 0, fmt.Errorf("creating task: %w", err)
	}

	return id, nil
}

func (s TaskService) UpdateTask(ctx context.Context, task domain.Task) (int, error) {
	id, err := s.repo.UpdateTask(ctx, task)
	if err != nil {
		return 0, fmt.Errorf("updating task: %w", err)
	}

	return id, nil
}

func (s TaskService) DeleteTask(ctx context.Context, id int) error {
	err := s.repo.DeleteTask(ctx, id)
	if err != nil {
		return fmt.Errorf("deleting task: %w", err)
	}

	return nil
}

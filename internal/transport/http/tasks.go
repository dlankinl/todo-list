package http

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
	"todo/internal/domain"
	"todo/internal/utils"
	"todo/models"
)

type TaskService interface {
	GetTasks(ctx context.Context) ([]domain.Task, error)
	GetTask(ctx context.Context, id int) (domain.Task, error)
	GetTasksByPriority(ctx context.Context, priority models.Priority) ([]domain.Task, error)
	CreateTask(ctx context.Context, task domain.Task) (int, error)
	UpdateTask(ctx context.Context, task domain.Task) (int, error)
	DeleteTask(ctx context.Context, id int) error
}

type Handler struct {
	service TaskService
}

func NewHandler(service TaskService) Handler {
	return Handler{service: service}
}

func (h Handler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.service.GetTasks(r.Context())
	if err != nil {
		log.Printf("error while getting tasks: %v", err)
		utils.ErrorJSON(w, errors.New("couldn't get tasks"), http.StatusBadRequest)
		return
	}

	utils.WriteJSON(
		w,
		http.StatusOK,
		utils.Envelope{
			"status": "ok",
			"tasks":  tasks,
		},
	)
}

func (h Handler) GetTasksByPriority(w http.ResponseWriter, r *http.Request) {
	priority := chi.URLParam(r, "priority")
	intPriority, err := strconv.Atoi(priority)
	if err != nil {
		log.Printf("error while parsing priority: %v", err)
		utils.ErrorJSON(w, errors.New("couldn't parse priority"), http.StatusBadRequest)
		return
	}

	tasks, err := h.service.GetTasksByPriority(r.Context(), models.Priority(intPriority))
	if err != nil {
		log.Printf("error while getting tasks by priority: %v", err)
		utils.ErrorJSON(w, errors.New("couldn't get tasks"), http.StatusBadRequest)
		return
	}

	utils.WriteJSON(
		w,
		http.StatusOK,
		utils.Envelope{
			"status":   "ok",
			"priority": intPriority,
			"tasks":    tasks,
		},
	)
}

func (h Handler) GetTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("error while parsing task id: %v", err)
		utils.ErrorJSON(w, errors.New("couldn't parse task id"), http.StatusBadRequest)
		return
	}

	task, err := h.service.GetTask(r.Context(), intID)
	if err != nil {
		log.Printf("error while getting task by id=%d: %v", intID, err)
		utils.ErrorJSON(w, errors.New("couldn't get task"), http.StatusBadRequest)
		return
	}

	utils.WriteJSON(
		w,
		http.StatusOK,
		utils.Envelope{
			"status": "ok",
			"id":     intID,
			"task":   task,
		},
	)
}

func (h Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var taskReq domain.Task
	err := json.NewDecoder(r.Body).Decode(&taskReq)
	if err != nil {
		log.Printf("error while decoding json: %v", err)
		utils.ErrorJSON(w, errors.New("couldn't decode json"), http.StatusBadRequest)
		return
	}

	id, err := h.service.CreateTask(r.Context(), taskReq)
	if err != nil {
		log.Printf("error while creating task: %v", err)
		utils.ErrorJSON(w, errors.New("couldn't create task"), http.StatusBadRequest)
		return
	}

	utils.WriteJSON(
		w,
		http.StatusOK,
		utils.Envelope{
			"status": "ok",
			"id":     id,
		},
	)
}

func (h Handler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	var taskReq domain.Task
	err := json.NewDecoder(r.Body).Decode(&taskReq)
	if err != nil {
		log.Printf("error while decoding json: %v", err)
		utils.ErrorJSON(w, errors.New("couldn't decode json"), http.StatusBadRequest)
		return
	}

	id, err := h.service.UpdateTask(r.Context(), taskReq)
	if err != nil {
		log.Printf("error while updating task id=%d: %v", taskReq.TaskID, err)
		utils.ErrorJSON(w, errors.New("couldn't update task"), http.StatusBadRequest)
		return
	}

	utils.WriteJSON(
		w,
		http.StatusOK,
		utils.Envelope{
			"status": "ok",
			"id":     id,
		},
	)
}

func (h Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("error while parsing task id: %s", err)
		utils.ErrorJSON(w, errors.New("couldn't parse task id"), http.StatusBadRequest)
		return
	}

	err = h.service.DeleteTask(r.Context(), intID)
	if err != nil {
		log.Printf("error while deleting task with id=%d: %v", intID, err)
		utils.ErrorJSON(w, errors.New("couldn't delete task"), http.StatusBadRequest)
		return
	}

	utils.WriteJSON(
		w,
		http.StatusOK,
		utils.Envelope{
			"status": "ok",
		},
	)
}

package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
	"todo/pkg/utils"
	"todo/services"
)

var task services.Task

// GET/tasks
func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	var task services.Task
	all, err := task.GetAllTasks()
	if err != nil {
		log.Fatal("error while getting tasks: ", err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"tasks": all})
}

// GET/task/{id}
func GetTaskByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		log.Fatalf("error while parsing task id (%s): %s", id, err)
		return
	}
	task, err := task.GetTaskByID(intID)
	if err != nil {
		log.Fatalf("error while getting task id (%d): %s", intID, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, task)
}

// POST/task
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var taskResp services.Task
	err := json.NewDecoder(r.Body).Decode(&taskResp)
	if err != nil {
		log.Print("error while adding task: ", err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, task)
	taskCreated, err := task.CreateTask(taskResp)
	if err != nil {
		log.Print("error while adding task: ", err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, taskCreated)
}

// PUT/task/{id}
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("error while parsing task id (%s): %s", id, err)
		return
	}
	var taskResp services.Task
	err = json.NewDecoder(r.Body).Decode(&taskResp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	utils.WriteJSON(w, http.StatusOK, taskResp)
	taskObj, err := task.UpdateTask(intID, taskResp)
	if err != nil {
		log.Fatalf("error while updating task (%d): %s", intID, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, taskObj)
}

// DELETE/task/{id}
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		log.Fatalf("error while parsing task id (%s): %s", id, err)
		return
	}
	err = task.DeleteTask(intID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	utils.WriteJSON(w, http.StatusOK, "succesful deletion")
}

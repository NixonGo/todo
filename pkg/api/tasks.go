package api

import (
	"net/http"

	"github.com/NixonGo/todo/pkg/db"
)

type TasksResponse struct {
	Tasks []*db.Task `json:"tasks"`
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := db.Tasks(taskLimit)
	if err != nil {
		writeError(w, http.StatusBadRequest, "cannot get task list")
		return
	}
	if tasks == nil {
		tasks = []*db.Task{}
	}

	writeJSON(w, http.StatusOK, TasksResponse{
		Tasks: tasks,
	})
}

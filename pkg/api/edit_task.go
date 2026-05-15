package api

import (
	"encoding/json"
	"net/http"

	"github.com/NixonGo/todo/pkg/db"
)

func editTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		writeError(w, http.StatusBadRequest, "cannot decode body")
		return
	}
	if task.ID == "" {
		writeError(w, http.StatusBadRequest, "No id")
		return
	}
	if task.Title == "" {
		writeError(w, http.StatusBadRequest, "No Title")
		return
	}

	err = checkDate(&task)
	if err != nil {
		writeError(w, http.StatusBadRequest, "cannot check date")
		return
	}
	err = db.UpdateTask(&task)
	if err != nil {
		writeError(w, http.StatusBadRequest, "cannot update task")
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{})

}

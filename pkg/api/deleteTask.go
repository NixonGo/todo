package api

import (
	"net/http"
	"todo/pkg/db"
)

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		writeError(w, http.StatusBadRequest, "cannot get id")
		return
	}
	err := db.DeleteTask(id)
	if err != nil {
		writeError(w, http.StatusBadRequest, "cannot delete task")
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{})
}

package api

import (
	"net/http"

	"github.com/NixonGo/todo/pkg/db"
)

func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		writeError(w, http.StatusBadRequest, "cannot get id")
		return
	}
	task, err := db.GetTask(id)
	if err != nil {
		writeError(w, http.StatusBadRequest, "cannot find task")
		return
	}
	writeJSON(w, http.StatusOK, task)
}

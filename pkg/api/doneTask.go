package api

import (
	"net/http"
	"time"
	"todo/pkg/db"
)

func doneTaskHandler(w http.ResponseWriter, r *http.Request) {
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

	if task.Repeat == "" {
		err = db.DeleteTask(id)
		if err != nil {
			writeError(w, http.StatusBadRequest, "cannot delete task")
			return
		}
	} else {
		next, err := NextDate(time.Now(), task.Date, task.Repeat)
		if err != nil {
			writeError(w, http.StatusBadRequest, "cannot calculate next date")
			return
		}
		err = db.UpdateDate(next, id)
		if err != nil {
			writeError(w, http.StatusBadRequest, "cannot update date")
			return
		}
	}
	writeJSON(w, http.StatusOK, map[string]string{})
}

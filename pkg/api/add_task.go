package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/NixonGo/todo/pkg/db"
)

const (
	dateLayout = "20060102"
	taskLimit  = 50
)

func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, status int, msg string) {
	err := writeJSON(w, status, map[string]string{
		"error": msg,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func checkDate(task *db.Task) error {
	now := time.Now()
	now = time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		0, 0, 0, 0,
		now.Location(),
	)

	if task.Date == "" {
		task.Date = now.Format(dateLayout)
	}

	date, err := time.Parse(dateLayout, task.Date)
	if err != nil {
		return errors.New("error parse date")
	}

	if now.After(date) && task.Repeat == "" {
		task.Date = now.Format(dateLayout)
	}

	if now.After(date) && task.Repeat != "" {
		newdate, err := NextDate(now, task.Date, task.Repeat)
		if err != nil {
			return err
		}

		task.Date = newdate
	}

	return nil
}

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		writeError(w, http.StatusBadRequest, "cannot decode body")
		return
	}

	if task.Title == "" {
		writeError(w, http.StatusBadRequest, "empty title")
		return
	}

	err = checkDate(&task)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := db.AddTask(&task)
	if err != nil {
		writeError(w, http.StatusBadRequest, "cannot add task")
		return
	}

	resp := map[string]string{
		"id": strconv.FormatInt(id, 10),
	}

	err = writeJSON(w, http.StatusOK, resp)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "cannot encode json")
		return
	}
}

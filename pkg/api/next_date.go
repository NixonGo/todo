package api

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func nextDateHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	nowStr := r.FormValue("now")
	dstart := r.FormValue("date")
	repeat := r.FormValue("repeat")

	var now time.Time
	if nowStr == "" {
		now = time.Now()
	} else {
		t, err := time.Parse(dateLayout, nowStr)
		if err != nil {
			http.Error(w, "bad now parameter", http.StatusBadRequest)
			return
		}
		now = t
	}

	result, err := NextDate(now, dstart, repeat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, result)
}

func NextDate(now time.Time, dstart string, repeat string) (string, error) {

	if repeat == "" {
		return "", errors.New("repeat is empty")
	}

	dstartTime, err := time.Parse(dateLayout, dstart)
	if err != nil {
		return "", err
	}

	repeatRule := strings.Split(repeat, " ")

	switch repeatRule[0] {
	case "y":
		dstartTime = dstartTime.AddDate(1, 0, 0)

		for !dstartTime.After(now) {
			dstartTime = dstartTime.AddDate(1, 0, 0)
		}

	case "d":
		if len(repeatRule) < 2 {
			return "", errors.New("missing argument for d")
		}

		drule, err := strconv.Atoi(repeatRule[1])
		if err != nil {
			return "", err
		}
		if drule <= 0 || drule > 400 {
			return "", errors.New("d argument out of range 1-400")
		}

		dstartTime = dstartTime.AddDate(0, 0, drule)

		for !dstartTime.After(now) {
			dstartTime = dstartTime.AddDate(0, 0, drule)
		}

	default:
		return "", errors.New("unsupported repeat rule")
	}

	return dstartTime.Format(dateLayout), nil
}

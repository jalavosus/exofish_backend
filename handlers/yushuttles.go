package handlers

import (
	"bytes"
	"encoding/json"
	"exofish-backend/yushuttles"
	"log"
	"net/http"
	"strconv"
)

func GetTimesHandler(w http.ResponseWriter, r *http.Request) {
	SetHeaders("GET", &w)

	direction, ok := r.URL.Query()["direction"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Please provide a direction as an int [`1` for uptown, `2` for downtown]"))

		return
	}

	bookingDate, ok := r.URL.Query()["date"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Please provide a date as string [ex. 2019-08-23]"))

		return
	}

	dirInt, _ := strconv.ParseInt(direction[0], 10, 0)
	directionInt := int(dirInt)

	availableTimes := yushuttles.GetShuttleTimes(directionInt, bookingDate[0])

	response := bytes.Buffer{}
	err := json.NewEncoder(&response).Encode(availableTimes)
	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(response.Bytes())
}

func BookShuttleHandler(w http.ResponseWriter, r *http.Request) {
	SetHeaders("POST", &w)

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	formData := BuildFormData(r.PostForm)

	htmlMessageBody, textMessageBody := yushuttles.LoadTemplate(formData)
	yushuttles.SendMail(formData["name"], formData["email"], htmlMessageBody, textMessageBody)

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("All good in the hood."))
}

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func getTimesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

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

	availableTimes := GetShuttleTimes(directionInt, bookingDate[0])

	response := bytes.Buffer{}
	_ = json.NewEncoder(&response).Encode(availableTimes)

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(response.Bytes())
}

func shuttlesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	formData := make(map[string]string)

	for k, v := range r.PostForm {
		formData[k] = v[0]
	}

	htmlMessageBody, textMessageBody := LoadTemplate(formData)
	SendMail(formData["name"], formData["email"], htmlMessageBody, textMessageBody)

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("All good in the hood."))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to rootland")
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/shuttles", shuttlesHandler)
	http.HandleFunc("/shuttleTimes", getTimesHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

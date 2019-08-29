package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func buildFormData(httpPostForm url.Values) map[string]string {
	formData := make(map[string]string)

	for k, v := range httpPostForm {
		formData[k] = v[0]
	}

	return formData
}

func setHeaders(method string, writer *http.ResponseWriter) {
	(*writer).Header().Set("Access-Control-Allow-Origin", "*")
	switch method {
	case "POST":
		(*writer).Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	case "GET":
		(*writer).Header().Set("Access-Control-Allow-Methods", "GET")
	}
}

func getTimesHandler(w http.ResponseWriter, r *http.Request) {
	setHeaders("GET", &w)

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
	err := json.NewEncoder(&response).Encode(availableTimes)
	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(response.Bytes())
}

func shuttlesHandler(w http.ResponseWriter, r *http.Request) {
	setHeaders("POST", &w)

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	formData := buildFormData(r.PostForm)

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

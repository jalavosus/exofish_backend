package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"crypto/tls"
	"encoding/base64"
	"net/http"
	"net/url"

	"encoding/json"
	"io/ioutil"
)

type TimesList struct {
	AvailableTimes []string `json:"available_times"`
}

const (
	WPAdminURL       string = "https://yushuttles.com/wp-admin/admin-ajax.php"
	UptownCategory   string = "4"
	DowntownCategory string = "5"
	UptownStaffID    string = "7958"
	DowntownStaffID  string = "7959"
)

var (
	RequestHeaders map[string]string
	HTTPClient     *http.Client
	FmtUser        string
)

func init() {
	HTTPClient = &http.Client{}
	RequestHeaders = map[string]string{
		"Accept":           "*/*",
		"Host":             "yushuttles.com",
		"Origin":           "https://yushuttles.com",
		"Referer":          "https://yushuttles.com/client-profile/?tab=appointment",
		"Sec-Fetch-Site":   "same-origin",
		"User-Agent":       "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.100 Safari/537.36",
		"X-Requested-With": "XMLHttpRequest",
		"Content-Type":     "application/x-www-form-urlencoded",
	}

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	FmtUser = fmt.Sprintf("this is an unused import")
}

func CreateTimelistForm(direction int, bookingDate string) map[string]string {

	date, _ := time.Parse(ISODateFormat, bookingDate)
	bookingDate = date.Format(AmericanDateFormat)

	formData := map[string]string{
		"action":      "ubp_book_step_2",
		"b_category":  "",
		"b_date":      bookingDate,
		"b_staff":     "",
		"b_location":  "",
		"template_id": "",
	}

	switch direction {
	case 1: // Going uptown
		formData["b_category"] = UptownCategory
		formData["b_staff"] = UptownStaffID
	case 2: // Going downtown
		formData["b_category"] = DowntownCategory
		formData["b_staff"] = DowntownStaffID
	}

	return formData
}

/**
 * Docstring here
 */
func GetShuttleTimes(direction int, bookingDate string) TimesList {
	timeListForm := CreateTimelistForm(direction, bookingDate)

	response := doHTTPRequest(WPAdminURL, timeListForm)

	strResponse := string(response)

	data, failed := base64Decode(strResponse)
	if failed {
		log.Println("halp")
	}

	jsonMap := jsonify(data)

	availableTimes := GetAvailableTimes(jsonMap["content"])

	tl := TimesList{
		AvailableTimes: availableTimes,
	}

	return tl
}

func doHTTPRequest(reqUrl string, params map[string]string) []byte {
	formData := url.Values{}

	for k, v := range params {
		formData.Set(k, v)
	}

	req, err := http.NewRequest("POST", reqUrl, strings.NewReader(formData.Encode()))
	if err != nil {
		log.Println(err)
	}

	for k, v := range RequestHeaders {
		req.Header.Set(k, v)
	}

	resp, err := HTTPClient.Do(req)
	if err != nil {
		log.Println(err)
	}

	response := loadResponse(resp)

	return response
}

func loadResponse(response *http.Response) []byte {
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}

	marshaledData, err := json.Marshal(body)
	if err != nil {
		log.Println(err)
	}

	return marshaledData
}

func jsonify(jsonStr string) map[string]string {
	jsonMap := make(map[string]string)

	err := json.Unmarshal([]byte(jsonStr), &jsonMap)
	if err != nil {
		log.Println(err)
	}

	return jsonMap
}

func base64Decode(str string) (string, bool) {
	str = strings.ReplaceAll(str, "\"", "")

	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		log.Println(err)
		return "", true
	}
	return string(data), false
}

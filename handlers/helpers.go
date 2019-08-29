package handlers

import (
	"net/http"
	"net/url"
)

func BuildFormData(httpPostForm url.Values) map[string]string {
	formData := make(map[string]string)

	for k, v := range httpPostForm {
		formData[k] = v[0]
	}

	return formData
}

/**
 * Takes an HTTP method (GET or POST), and sets the
 * Access-Control-Allow-Methods headers of `writer` accordingly.
 */
func SetHeaders(method string, writer *http.ResponseWriter) {
	(*writer).Header().Set("Access-Control-Allow-Origin", "*")
	switch method {
	case "POST":
		(*writer).Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	case "GET":
		(*writer).Header().Set("Access-Control-Allow-Methods", "GET")
	}
}

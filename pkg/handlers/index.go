package handlers

import (
	"net/http"

	"jilt.com/m/pkg/models"
)

func Index(writer http.ResponseWriter, request *http.Request) {
	topics, err := models.GetAllTopics()
	if err != nil {
		http.Redirect(writer, request, "/err?msg=Cannot get topics", http.StatusTemporaryRedirect)
	} else {
		_, err := session(writer, request)
		if err != nil {
			generateHTML(writer, topics, "layout", "navbar", "index")
		} else {
			generateHTML(writer, topics, "layout", "auth.navbar", "index")
		}
	}
}

func Err(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	_, err := session(writer, request)
	if err != nil {
		generateHTML(writer, vals.Get("msg"), "layout", "navbar", "error")
	} else {
		generateHTML(writer, vals.Get("msg"), "layout", "auth.navbar", "error")
	}
}

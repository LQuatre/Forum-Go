package handlers

import (
	"fmt"
	"net/http"

	"jilt.com/m/pkg/models"
)

func Index(writer http.ResponseWriter, request *http.Request) {
	categories, err := models.GetAllCategories()
	fmt.Println(categories)
	if err != nil {
		http.Redirect(writer, request, "/err?msg=Cannot get topics", http.StatusTemporaryRedirect)
	} else {
		_, err := session(writer, request)
		if err != nil {
			fmt.Println("No session")
			generateHTML(writer, nil, "layout", "navbar", "index")
		} else {
			fmt.Println("Session")
			generateHTML(writer, categories, "layout", "auth.navbar", "auth.index", "categories")
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

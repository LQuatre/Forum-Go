package handlers

import (
	"fmt"
	"net/http"

	"jilt.com/m/pkg/models"
)

func Index(writer http.ResponseWriter, request *http.Request) {
	categories, err := models.GetAllCategories()
	topics, err := models.GetAllTopics()
	for i, _ := range categories {
		for j, _ := range topics {
			if topics[j].CategoryUuId == categories[i].Uuid {
				categories[i].Topics = append(categories[i].Topics, topics[j])
			}
		}
	}

	fmt.Println(categories)
	if err != nil {
		http.Redirect(writer, request, "/err?msg=Cannot get topics", http.StatusTemporaryRedirect)
	} else {
		_, err := session(writer, request)
		if err != nil {
			generateHTML(writer, nil, "layout", "navbar", "index")
		} else {
			generateHTML(writer, categories, "layout", "auth.navbar", "auth.index")
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

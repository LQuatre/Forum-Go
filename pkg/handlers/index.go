package handlers

import (
	"net/http"

	"jilt.com/m/pkg/models"
)

type IndexData struct {
	Categories []models.Category
	User       models.User
}

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

	if err != nil {
		http.Redirect(writer, request, "/err?msg=Cannot get topics", http.StatusTemporaryRedirect)
	} else {
		sess, err := session(writer, request)
		if err != nil {
			user := models.User{
				Name:  "Guest",
				Email: "",
			}
			indexData := IndexData{
				Categories: categories,
				User:       user,
			}
			generateHTML(writer, &indexData, "layout", "navbar", "index")
		} else {
			user, err := sess.User()
			indexData := IndexData{
				Categories: categories,
				User:       user,
			}
			if err != nil {
				danger(err, "Cannot get user from session")
			}
			if user.IsAdmin {
				generateHTML(writer, &indexData, "layout", "admin.navbar", "index")
			} else {
				generateHTML(writer, &indexData, "layout", "auth.navbar", "index")
			}
		}
	}
}

func Help(writer http.ResponseWriter, request *http.Request) {
	_, err := session(writer, request)
	if err != nil {
		generateHTML(writer, nil, "layout", "navbar", "help")
	} else {
		generateHTML(writer, nil, "layout", "auth.navbar", "help")
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

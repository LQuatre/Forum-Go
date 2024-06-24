package handlers

import (
	"net/http"

	"jilt.com/m/pkg/models"
)

// GET /topics
func Topics(writer http.ResponseWriter, request *http.Request) {
	// redirect to categories
	http.Redirect(writer, request, "/categories", 302)
}

// GET /topics/new
func NewTopic(writer http.ResponseWriter, request *http.Request) {
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		categories, err := models.Categories()
		if err != nil {
			danger(err, "Cannot get categories")
		}
		generateHTML(writer, &categories, "layout", "auth.navbar", "new.topic")
	}
}

// POST /topics/create
func CreateTopic(writer http.ResponseWriter, request *http.Request) {
	sess, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			danger(err, "Cannot parse form")
		}
		user, err := sess.User()
		if err != nil {
			danger(err, "Cannot get user from session")
		}
		name := request.PostFormValue("name")
		categoryuuid := request.PostFormValue("cat-uuid")
		if _, err := user.CreateTopic(name, categoryuuid); err != nil {
			danger(err, "Cannot create topic")
		}
		http.Redirect(writer, request, "/categories/category?uuid="+categoryuuid, 302)
	}
}

// POST /topics/delete
func DeleteTopic(writer http.ResponseWriter, request *http.Request) {
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			danger(err, "Cannot parse form")
		}
		uuid := request.PostFormValue("uuid")
		if err := models.DeleteTopic(uuid); err != nil {
			danger(err, "Cannot delete topic")
		}
		http.Redirect(writer, request, "/topics", 302)
	}
}

// GET /topics/topic
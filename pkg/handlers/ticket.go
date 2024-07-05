package handlers

import (
	"net/http"

	"jilt.com/m/pkg/models"
)

func Chat(writer http.ResponseWriter, request *http.Request) {
	sess, err := session(writer, request)
	if err != nil {
		user := models.User{
			Name:  "Guest",
			Email: "",
		}
		generateHTML(writer, user, "layout", "navbar", "chat")
	} else {
		user, err := sess.User()
		if err != nil {
			danger(err, "Cannot get user from session")
		}
		generateHTML(writer, user, "layout", "auth.navbar", "chat")
	}
}
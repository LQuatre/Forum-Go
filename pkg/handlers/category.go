package handlers

import (
	"net/http"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"jilt.com/m/pkg/models"
)

func Categories(writer http.ResponseWriter, request *http.Request) {
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		categories, err := models.Categories()
		if err != nil {
			danger(err, "Cannot get categories")
		}
		generateHTML(writer, &categories, "layout", "auth.navbar", "categories")
	}
}

// GET /category/new
func NewCategory(writer http.ResponseWriter, request *http.Request) {
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		generateHTML(writer, nil, "layout", "auth.navbar", "new.category")
	}
}

// POST /category/create
func CreateCategory(writer http.ResponseWriter, request *http.Request) {
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
		if _, err := user.CreateCategory(name); err != nil {
			danger(err, "Cannot create category")
		}
		http.Redirect(writer, request, "/", 302)
	}
}

// GET /category/read
func GoCategory(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	uuid := vals.Get("id")
	category, err := models.CategoryByUUID(uuid)
	if err != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "category_not_found",
		})
		errorMessage(writer, request, msg)
	} else {
		_, err := session(writer, request)
		if err != nil {
			generateHTML(writer, &category, "layout", "public.navbar", "public.category")
		} else {
			generateHTML(writer, &category, "layout", "auth.navbar", "auth.category")
		}
	}
}

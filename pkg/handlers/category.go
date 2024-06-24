package handlers

import (
	"fmt"
	"net/http"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"jilt.com/m/pkg/models"
)

// GET /categories
func Categories(writer http.ResponseWriter, request *http.Request) {
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		categories, err := models.Categories()
		if err != nil {
			danger(err, "Cannot get categories")
		}
		topics, err := models.Topics()
		for i, _ := range categories {
			for j, _ := range topics {
				if topics[j].CategoryUuId == categories[i].Uuid {
					categories[i].Topics = append(categories[i].Topics, topics[j])
				}
			}
		}
		if err != nil {
			danger(err, "Cannot get topics")
		}
		generateHTML(writer, &categories, "layout", "auth.navbar", "auth.categories")
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
		http.Redirect(writer, request, "/categories", 302)
	}
}

// POST /categories/delete
func DeleteCategory(writer http.ResponseWriter, request *http.Request) {
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		vals := request.URL.Query()
		uuid := vals.Get("uuid")
		fmt.Println("UUID: ", uuid)
		category, err := models.CategoryByUUID(uuid)
		if err != nil {
			danger(err, "Cannot delete category")
		}
		if err := category.Delete(); err != nil {
			danger(err, "Cannot delete category")
		}
		http.Redirect(writer, request, "/categories", 302)
	}
}

// GET /categories/category
func GoCategory(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	uuid := vals.Get("uuid")
	fmt.Println("UUID: ", uuid)
	category, err := models.CategoryByUUID(uuid)
	if err != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "category_not_found",
		})
		errorMessage(writer, request, msg)
	} else {
		topics, err := models.TopicsFromCategoryUUID(uuid)
		if err != nil {
			msg := localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID: "cannot_get_topics",
			})
			errorMessage(writer, request, msg)
		}
		category.Topics = topics
		_, err = session(writer, request)
		if err != nil {
			generateHTML(writer, &category, "layout", "public.navbar", "public.category")
		} else {
			generateHTML(writer, &category, "auth.layout", "auth.navbar", "auth.category")
		}
	}
}

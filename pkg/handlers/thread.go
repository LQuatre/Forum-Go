package handlers

import (
	"net/http"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"jilt.com/m/pkg/models"
)

// GET /threads/new
func NewThread(writer http.ResponseWriter, request *http.Request) {
    _, err := session(writer, request)
    if err != nil {
        http.Redirect(writer, request, "/login", 302)
    } else {
        generateHTML(writer, nil, "layout", "auth.navbar", "new.thread")
    }
}

// POST /thread/create
func CreateThread(writer http.ResponseWriter, request *http.Request) {
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
        topicuuid, name, desc := request.PostFormValue("topic-uuid"), request.PostFormValue("name"), request.PostFormValue("desc");
        if err != nil {
            danger(err, "Invalid topic UUID")
            return
        }
        if _, err := user.CreateThread(topicuuid, name, desc); err != nil {
            danger(err, "Cannot create thread")
        }
        http.Redirect(writer, request, "/", 302)
    }
}

// GET /thread/read
func ReadThread(writer http.ResponseWriter, request *http.Request) {
    vals := request.URL.Query()
    uuid := vals.Get("uuid")
    thread, err := models.ThreadByUUID(uuid)
    if err != nil {
        msg := localizer.MustLocalize(&i18n.LocalizeConfig{
            MessageID: "thread_not_found",
        })
        errorMessage(writer, request, msg)
    } else {
        comments, err := models.GetCommentsByThreadUUID(uuid)
        if err != nil {
            danger(err, "Cannot get comments")
        }
        thread.Comments = comments
        for i, _ := range comments {
            user, err := models.UserByUUID(comments[i].UserUuId)
            if err != nil {
                danger(err, "Cannot get user")
            }
            comments[i].Author = user
            likes, err := models.GetLikesByPostUUID(comments[i].Uuid)
            if err != nil {
                danger(err, "Cannot get likes")
            }
            countLikes := 0
            for k, _ := range likes {
                likes[k].Value = 1
                if likes[k].Value == 1 {
                    countLikes++
                }
            }
            comments[i].Likes = countLikes
            dislikes, err := models.GetDislikesByPostUUID(comments[i].Uuid)
            if err != nil {
                danger(err, "Cannot get dislikes")
            }
            countDislikes := 0
            for k, _ := range dislikes {
                dislikes[k].Value = -1
                if dislikes[k].Value == -1 {
                    countDislikes++
                }
            }
            comments[i].Dislikes = countDislikes
        }
        sess, err := session(writer, request)
        if err != nil {
            generateHTML(writer, &thread, "layout", "navbar", "thread")
        } else {
            user, err := sess.User()
            if err != nil {
                danger(err, "Cannot get user from session")
            }
            if user.IsAdmin {
                generateHTML(writer, &thread, "layout", "admin.navbar", "thread")
            } else {
                generateHTML(writer, &thread, "layout", "auth.navbar", "thread")
            }
        }
    }
}

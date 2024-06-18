package handlers

import (
	"net/http"
)

func Index(writer http.ResponseWriter, request *http.Request) {
	generateHTML(writer, "", "index")
}

// func Err(writer http.ResponseWriter, request *http.Request)  {
//     vals := request.URL.Query()
//     _, err := session(writer, request)
//     if err != nil {
//         generateHTML(writer, vals.Get("msg"), "layout", "navbar", "error")
//     } else {
//         generateHTML(writer, vals.Get("msg"), "layout", "auth.navbar", "error")
//     }
// }
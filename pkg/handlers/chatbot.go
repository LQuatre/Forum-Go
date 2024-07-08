package handlers

import (
	"fmt"
	"net/http"

	"jilt.com/m/pkg/models"
)

func ChatBotCreateTicket(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("ChatBotCreateTicket")
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			danger(err, "Cannot parse form")
		}
		fmt.Printf("name: %s\n", request.PostFormValue("name"))
		fmt.Printf("user_uuid: %s\n", request.PostFormValue("user_uuid"))
		fmt.Printf("description: %s\n", request.PostFormValue("description"))
		models.CreateTicket(request.PostFormValue("name"), request.PostFormValue("user_uuid"), request.PostFormValue("description"))
	}
}

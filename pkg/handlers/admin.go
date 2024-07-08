package handlers

import (
	"fmt"
	"net/http"

	"jilt.com/m/pkg/models"
)

func Admin(writer http.ResponseWriter, request *http.Request) {
	session, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		user, err := session.User()
		if err != nil {
			danger(err, "Cannot get user from session")
		}
		if user.IsAdmin != true {
			http.Redirect(writer, request, "/", 302)
		}
		admin, err := models.Admin()
		if err != nil {
			danger(err, "Cannot get admin info")
		}
		generateHTML(writer, admin, "layout", "admin.navbar", "admin")
	}
}

func AdminUpdate(writer http.ResponseWriter, request *http.Request) {
	session, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		user, err := session.User()
		if err != nil {
			danger(err, "Cannot get user from session")
		}
		if user.IsAdmin != true {
			http.Redirect(writer, request, "/", 302)
		}
		err = request.ParseForm()
		if err != nil {
			danger(err, "Cannot parse form")
		}
		admin, err := models.UserByUUID(request.PostFormValue("admin-uuid"))
		if err != nil {
			danger(err, "Cannot get admin info")
		}
		if request.PostFormValue("name") != "" {
			admin.Name = request.PostFormValue("name")
		}
		if request.PostFormValue("email") != "" {
			admin.Email = request.PostFormValue("email")
		}
		if err := admin.Update(); err != nil {
			danger(err, "Cannot update admin info")
		}
		success("Admin info updated", "Admin info updated successfully")
		http.Redirect(writer, request, "/admin", 302)
	}
}

func AdminUpdate2(writer http.ResponseWriter, request *http.Request) {
	session, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		user, err := session.User()
		if err != nil {
			danger(err, "Cannot get user from session")
		}
		if user.IsAdmin != true {
			http.Redirect(writer, request, "/", 302)
		}
		err = request.ParseForm()
		if err != nil {
			danger(err, "Cannot parse form")
		}
		admin, err := models.UserByUUID(request.PostFormValue("admin-uuid"))
		if err != nil {
			danger(err, "Cannot get admin info")
		}
		admin.IsAdmin = !admin.IsAdmin
		if err := admin.Update2(); err != nil {
			danger(err, "Cannot update admin info")
		}
		success("Admin info updated", "Admin info updated successfully")
		http.Redirect(writer, request, "/admin", 302)
	}
}

func AdminCloseASession(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		danger(err, "Cannot parse form")
	}
	session := models.Session{Uuid: request.PostFormValue("session-uuid")}
	err = session.DeleteByUUID()
	if err != nil {
		danger(err, "Cannot delete session")
	}
	success("Session deleted", "Session deleted successfully")
	http.Redirect(writer, request, "/admin", 302)
}

type AdminHelp struct {
	Tickets     *[]models.Ticket
	Ticket      *models.Ticket
	Replies     *[]models.MessageTicket
	DefinedName string
	DefinedUUID string
}

func AdminHelps(writer http.ResponseWriter, request *http.Request) {
	session, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		user, err := session.User()
		if err != nil {
			danger(err, "Cannot get user from session")
		}
		if user.IsAdmin != true {
			http.Redirect(writer, request, "/", 302)
		}
		tickets, err := models.Tickets()
		if err != nil {
			danger(err, "Cannot get helps")
		}
		adminHelps := AdminHelp{
			Tickets:     &tickets,
			Ticket:      nil,
			Replies:     nil,
			DefinedName: "",
			DefinedUUID: "",
		}
		if request.URL.Query().Get("uuid") != "" {
			ticket, err := models.TicketByUUID(request.URL.Query().Get("uuid"))
			if err != nil {
				danger(err, "Cannot get ticket")
			}
			adminHelps.Ticket = &ticket
			replies, err := models.MessageTicketsByTicketUUID(ticket.Uuid)
			if err != nil {
				danger(err, "Cannot get replies")
			}
			adminHelps.Replies = &replies
			adminHelps.DefinedName = ticket.Name
			adminHelps.DefinedUUID = ticket.UserUuId
			fmt.Println(adminHelps.DefinedUUID)
		}
		generateHTML(writer, &adminHelps, "layout", "admin.navbar", "chat")
	}
}

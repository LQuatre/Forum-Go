package models

import "time"

type Ticket struct {
	Uuid         string
	Name         string
	Description  string
	UserUuId     string
	CreatedAt    time.Time
	Message      []MessageTicket
}

func Tickets() (tickets []Ticket, err error) {
	rows, err := Db.Query("SELECT uuid, name, user_uuid, created_at FROM tickets")
	if err != nil {
		return
	}
	for rows.Next() {
		ticket := Ticket{}
		if err = rows.Scan(&ticket.Uuid, &ticket.Name, &ticket.UserUuId, &ticket.CreatedAt); err != nil {
			return
		}
		tickets = append(tickets, ticket)
	}
	rows.Close()
	return
}

// Create a new ticket in a category
func CreateTicket(name string, userUuId string) (ticket Ticket, err error) {
	statement := "INSERT INTO tickets (uuid, name, user_uuid, created_at) VALUES (?, ?, ?, ?)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	uuid := createUUID()
	_, err = stmt.Exec(uuid, name, userUuId, time.Now())
	if err != nil {
		return
	}

	ticket = Ticket{
		Uuid:         uuid,
		Name:         name,
		UserUuId:     userUuId,
		CreatedAt:    time.Now(),
	}
	return
}

// Get all tickets in a category
func TicketsByUser(userUuId string) (tickets []Ticket, err error) {
	rows, err := Db.Query("SELECT uuid, name, user_uuid, created_at FROM tickets WHERE user_uuid = ?", userUuId)
	if err != nil {
		return
	}
	for rows.Next() {
		ticket := Ticket{}
		if err = rows.Scan(&ticket.Uuid, &ticket.Name, &ticket.UserUuId, &ticket.CreatedAt); err != nil {
			return
		}
		tickets = append(tickets, ticket)
	}
	rows.Close()
	return
}

// Get a ticket by UUID
func TicketByUUID(uuid string) (ticket Ticket, err error) {
	ticket = Ticket{}
	err = Db.QueryRow("SELECT uuid, name, user_uuid, created_at FROM tickets WHERE uuid = ?", uuid).
		Scan(&ticket.Uuid, &ticket.Name, &ticket.UserUuId, &ticket.CreatedAt)
	return
}

// Get the user who started this ticket
func (ticket *Ticket) User() (user User) {
	user = User{}
	Db.QueryRow("SELECT uuid, name, email, created_at FROM users WHERE uuid = ?", ticket.UserUuId).
		Scan(&user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	return
}

func (ticket *Ticket) Messages() (messages []MessageTicket, err error) {
	rows, err := Db.Query("SELECT uuid, body, user_uuid, ticket_uuid, created_at FROM messages WHERE ticket_uuid = ?", ticket.Uuid)
	if err != nil {
		return
	}
	for rows.Next() {
		message := MessageTicket{}
		if err = rows.Scan(&message.Uuid, &message.Body, &message.UserUuId, &message.TicketUuId, &message.CreatedAt); err != nil {
			return
		}
		messages = append(messages, message)
	}
	rows.Close()
	return
}

func (ticket *Ticket) CreateMessage(body string, userUuId string) (message MessageTicket, err error) {
	statement := "INSERT INTO messages (uuid, body, user_uuid, ticket_uuid, created_at) VALUES (?, ?, ?, ?, ?)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	uuid := createUUID()
	_, err = stmt.Exec(uuid, body, userUuId, ticket.Uuid, time.Now())
	if err != nil {
		return
	}

	message = MessageTicket{
		Uuid:         uuid,
		Body:         body,
		UserUuId:     userUuId,
		TicketUuId:   ticket.Uuid,
		CreatedAt:    time.Now(),
	}
	return
}
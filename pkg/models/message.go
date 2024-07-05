package models

import "time"

type MessageTicket struct {
	Uuid      string
	Body      string
	UserUuId  string
	TicketUuId string
	CreatedAt time.Time
}

func Messages() (messages []MessageTicket, err error) {
	rows, err := Db.Query("SELECT uuid, body, user_uuid, ticket_uuid, created_at FROM messages")
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

// Create a new message in a ticket
func CreateMessage(body string, userUuId string, ticketUuId string) (message MessageTicket, err error) {
	statement := "INSERT INTO messages (uuid, body, user_uuid, ticket_uuid, created_at) VALUES (?, ?, ?, ?, ?)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	uuid := createUUID()
	_, err = stmt.Exec(uuid, body, userUuId, ticketUuId, time.Now())
	if err != nil {
		return
	}

	message = MessageTicket{
		Uuid:         uuid,
		Body:         body,
		UserUuId:     userUuId,
		TicketUuId:   ticketUuId,
		CreatedAt:    time.Now(),
	}
	return
}

// Get all messages in a ticket
func MessagesByTicket(ticketUuId string) (messages []MessageTicket, err error) {
	rows, err := Db.Query("SELECT uuid, body, user_uuid, ticket_uuid, created_at FROM messages WHERE ticket_uuid = ?", ticketUuId)
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

// Get a message by UUID
func MessageByUUID(uuid string) (message MessageTicket, err error) {
	message = MessageTicket{}
	err = Db.QueryRow("SELECT uuid, body, user_uuid, ticket_uuid, created_at FROM messages WHERE uuid = ?", uuid).
		Scan(&message.Uuid, &message.Body, &message.UserUuId, &message.TicketUuId, &message.CreatedAt)
	return
}

// Get the user who started this message
func (message *MessageTicket) User() (user User, err error) {
	user = User{}
	err = Db.QueryRow("SELECT uuid, name, email, created_at, isAdmin FROM users WHERE uuid = ?", message.UserUuId).
		Scan(&user.Uuid, &user.Name, &user.Email, &user.CreatedAt, &user.IsAdmin)
	return
}
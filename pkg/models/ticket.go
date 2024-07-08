package models

import (
	"fmt"
	"time"
)

type Ticket struct {
	Uuid        string
	Name        string
	Description string
	UserUuId    string
	CreatedAt   time.Time
	Message     []MessageTicket
}

func Tickets() (tickets []Ticket, err error) {
	rows, err := Db.Query("SELECT uuid, name, desc, user_uuid, created_at FROM tickets")
	if err != nil {
		return
	}
	for rows.Next() {
		ticket := Ticket{}
		if err = rows.Scan(&ticket.Uuid, &ticket.Name, &ticket.Description, &ticket.UserUuId, &ticket.CreatedAt); err != nil {
			return
		}
		tickets = append(tickets, ticket)
	}
	rows.Close()
	return
}

// Create a new ticket in a category
func CreateTicket(name string, userUuId string, description string) (ticket Ticket, err error) {
	fmt.Println("CreateTicket")
	fmt.Printf("name: %s, description: %s, userUuId: %s\n", name, description, userUuId)
	statement := "INSERT INTO tickets (uuid, name, desc, user_uuid, created_at) VALUES (?, ?, ?, ?, ?)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	uuid := createUUID()
	_, err = stmt.Exec(uuid, name, description, userUuId, time.Now())
	if err != nil {
		return
	}

	ticket = Ticket{
		Uuid:        uuid,
		Name:        name,
		Description: description,
		UserUuId:    userUuId,
		CreatedAt:   time.Now(),
	}
	return
}

// Get all tickets in a category
func TicketsByUser(userUuId string) (tickets []Ticket, err error) {
	rows, err := Db.Query("SELECT uuid, name, desc, user_uuid, created_at FROM tickets WHERE user_uuid = ?", userUuId)
	if err != nil {
		return
	}
	for rows.Next() {
		ticket := Ticket{}
		if err = rows.Scan(&ticket.Uuid, &ticket.Name, &ticket.Description, &ticket.UserUuId, &ticket.CreatedAt); err != nil {
			return
		}
		tickets = append(tickets, ticket)
	}
	rows.Close()
	return
}

func TicketByUUID(uuid string) (ticket Ticket, err error) {
	ticket = Ticket{}
	err = Db.QueryRow("SELECT uuid, name, desc, user_uuid, created_at FROM tickets WHERE uuid = ?", uuid).
		Scan(&ticket.Uuid, &ticket.Name, &ticket.Description, &ticket.UserUuId, &ticket.CreatedAt)
	return
}

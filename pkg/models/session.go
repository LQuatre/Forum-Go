package models

import (
	"time"
)

type Session struct {
	Uuid      string
	Email     string
	UserUuId  string
	CreatedAt time.Time
	isAdmin   bool
}

// Check if session is valid in the database
func (session *Session) Check() (valid bool, err error) {
	err = Db.QueryRow("SELECT uuid, email, user_uuid, created_at, isAdmin FROM sessions WHERE uuid = ?", session.Uuid).
		Scan(&session.Uuid, &session.Email, &session.UserUuId, &session.CreatedAt, &session.isAdmin)
	if err != nil {
		valid = false
		return
	}
	if session.UserUuId != "" {
		valid = true
	}
	return
}

// Delete session from database
func (session *Session) DeleteByUUID() (err error) {
	statement := "DELETE FROM sessions WHERE uuid = ?"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(session.Uuid)
	return
}

// Get the user from the session
func (session *Session) User() (user User, err error) {
	user = User{}
	err = Db.QueryRow("SELECT uuid, name, email, created_at, isAdmin FROM users WHERE uuid = ?", session.UserUuId).
		Scan(&user.Uuid, &user.Name, &user.Email, &user.CreatedAt, &user.IsAdmin)
	return
}

func GetAllSessions() (sessions []Session, err error) {
	rows, err := Db.Query("SELECT uuid, email, user_uuid, created_at FROM sessions")
	if err != nil {
		return
	}
	for rows.Next() {
		session := Session{}
		if err = rows.Scan(&session.Uuid, &session.Email, &session.UserUuId, &session.CreatedAt); err != nil {
			return
		}
		sessions = append(sessions, session)
	}
	rows.Close()
	return
}

// Delete all sessions from database
func SessionDeleteAll() (err error) {
	statement := "DELETE FROM sessions"
	_, err = Db.Exec(statement)
	return
}

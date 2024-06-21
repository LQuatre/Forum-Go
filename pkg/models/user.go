package models

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	Id        int
	Uuid      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

func (user *User) Create() (err error) {
	// Check if name or email already exist
	existingUser, err := UserByName(user.Name)
	if err != nil {
		fmt.Printf("name not used 1 %v\n", err)
	}
	if existingUser.Id != 0 {
		fmt.Println("name already used 2")
		return errors.New("name already used")
	}

	existingUser, err = UserByEmail(user.Email)
	if err != nil {
		fmt.Printf("email not used 1 %v\n", err)
	}
	if existingUser.Id != 0 {
		fmt.Println("email already used 2")
		return errors.New("email already used")
	}

	statement := "INSERT INTO users (uuid, name, email, password, created_at) VALUES (?, ?, ?, ?, ?)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	uuid := createUUID()
	_, err = stmt.Exec(uuid, user.Name, user.Email, Encrypt(user.Password), time.Now())
	return
}

func UserByEmail(email string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE email = ?", email).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return user, err
}

func UserByName(name string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE name = ?", name).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return user, err
}

func UserByUUID(uuid string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE uuid = ?", uuid).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return user, err
}

func (user *User) Update() (err error) {
	statement := "UPDATE users SET name = ?, email = ? WHERE id = ?"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Email, user.Id)
	return err
}

func (user *User) Delete() (err error) {
	statement := "DELETE FROM users WHERE id = ?"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id)
	return
}

func (user *User) CreateSession() (session Session, err error) {
	statement := "INSERT INTO sessions (uuid, email, user_id, created_at) VALUES (?, ?, ?, ?)"
	stmtin, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmtin.Close()

	uuid := createUUID()
	stmtin.Exec(uuid, user.Email, user.Id, time.Now())

	stmtout, err := Db.Prepare("SELECT id, uuid, email, user_id, created_at FROM sessions WHERE uuid = ?")
	if err != nil {
		return
	}
	defer stmtout.Close()
	err = stmtout.QueryRow(uuid).Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	return
}

func (user *User) Session() (session Session, err error) {
	session = Session{}
	err = Db.QueryRow("SELECT id, uuid, email, user_id, created_at FROM sessions WHERE user_id = ?", user.Id).
		Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	return
}

func (user *User) CreateThread(topicId int, title string) (thread Thread, err error) {
	statement := "INSERT INTO threads (uuid, topic_id, user_id, title, created_at) VALUES (?, ?, ?, ?, ?)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	uuid := createUUID()
	_, err = stmt.Exec(uuid, topicId, user.Id, title, time.Now())
	if err != nil {
		return
	}

	stmt.QueryRow(uuid).Scan(&thread.Id, &thread.Uuid, &thread.TopicId, &thread.UserId, &thread.Title, &thread.CreatedAt)
	return
<<<<<<< HEAD
}

func (user *User) CreateCategory(name string) (category Category, err error) {
	statement := "INSERT INTO categories (uuid, name, created_at) VALUES (?, ?, ?)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	uuid := createUUID()
	_, err = stmt.Exec(uuid, name, time.Now())
	if err != nil {
		return
	}

	stmt.QueryRow(uuid).Scan(&category.Id, &category.Uuid, &category.Name, &category.CreatedAt)
	return
=======
>>>>>>> 8fe4e452f5f0b09a8ecbf7855b39ea78c56d5af3
}

func (user *User) CreatePost(thread *Thread, body string) (*Post, error) {
	statement := "INSERT INTO posts (uuid, body, user_id, thread_id, created_at) VALUES (?, ?, ?, ?, ?)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	uuid := createUUID()
	_, err = stmt.Exec(uuid, body, user.Id, thread.Id, time.Now())
	if err != nil {
		return nil, err
	}

	post := &Post{}
	stmt.QueryRow(uuid).Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt)
	return post, nil
}

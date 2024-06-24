package models

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
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
	if existingUser.Name != "" {
		return errors.New("Name already used")
	}

	existingUser, err = UserByEmail(user.Email)
	if err != nil {
		fmt.Printf("email not used 1 %v\n", err)
	}
	if existingUser.Email != "" {
		return errors.New("Email already used")
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
	err = Db.QueryRow("SELECT uuid, name, email, password, created_at FROM users WHERE email = ?", email).
		Scan(&user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return user, err
}

func UserByName(name string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("SELECT uuid, name, email, password, created_at FROM users WHERE name = ?", name).
		Scan(&user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return user, err
}

func UserByUUID(uuid string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("SELECT uuid, name, email, password, created_at FROM users WHERE uuid = ?", uuid).
		Scan(&user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return user, err
}

func (user *User) Update() (err error) {
	statement := "UPDATE users SET name = ?, email = ? WHERE Uuid = ?"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Email)
	return err
}

func (user *User) Delete() (err error) {
	statement := "DELETE FROM users WHERE Uuid = ?"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Uuid)
	return
}

func (user *User) CreateSession() (session Session, err error) {
	statement := "INSERT INTO sessions (uuid, email, user_uuid, created_at) VALUES (?, ?, ?, ?)"
	stmtin, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmtin.Close()

	uuid := createUUID()
	stmtin.Exec(uuid, user.Email, user.Uuid, time.Now())

	stmtout, err := Db.Prepare("SELECT uuid, email, created_at FROM sessions WHERE uuid = ?")
	if err != nil {
		return
	}
	defer stmtout.Close()
	err = stmtout.QueryRow(uuid).Scan(&session.Uuid, &session.Email, &session.CreatedAt)
	return
}

func (user *User) Session() (session Session, err error) {
	session = Session{}
	err = Db.QueryRow("SELECT uuid, email, created_at FROM sessions WHERE user_Uuid = ?", user.Uuid).
		Scan(&session.Uuid, &session.Email, &session.CreatedAt)
	return
}

func (user *User) CreateThread(topicUuId int, title string) (thread Thread, err error) {
	statement := "INSERT INTO threads (uuid, topic_uuid, user_uuid, title, created_at) VALUES (?, ?, ?, ?, ?)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	uuid := createUUID()
	_, err = stmt.Exec(uuid, topicUuId, user.Uuid, title, time.Now())
	if err != nil {
		return
	}

	stmt.QueryRow(uuid).Scan(&thread.Uuid, &thread.TopicUuId, &thread.UserUuId, &thread.Title, &thread.CreatedAt)
	return
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

	stmt.QueryRow(uuid).Scan(&category.Uuid, &category.Name, &category.CreatedAt)
	return
}

func (user *User) CreateTopic(name, categoryuuid string) (topic Topic, err error) {
	statement := "INSERT INTO topics (uuid, name, category_uuid, created_at) VALUES (?, ?, ?, ?)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	uuid := createUUID()
	_, err = stmt.Exec(uuid, name, categoryuuid, time.Now())
	if err != nil {
		return
	}

	stmt.QueryRow(uuid).Scan(&topic.Uuid, &topic.Name, &topic.CategoryUuId, &topic.CreatedAt)
	return
}


func (user *User) CreatePost(thread *Thread, body string) (*Post, error) {
	statement := "INSERT INTO posts (uuid, body, user_uuid, thread_uuid, created_at) VALUES (?, ?, ?, ?, ?)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	uuid := createUUID()
	_, err = stmt.Exec(uuid, body, user.Uuid, thread.Uuid, time.Now())
	if err != nil {
		return nil, err
	}

	post := &Post{}
	stmt.QueryRow(uuid).Scan(&post.Uuid, &post.Body, &post.UserUuId, &post.ThreadUuId, &post.CreatedAt)
	return post, nil
}

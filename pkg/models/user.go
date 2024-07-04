package models

import (
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Uuid      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	IsAdmin   bool
	DiscordID string
	GoogleID  string
	FacebookID string
	GithubID string
}

func (user *User) Create() (err error) {
	// Check if name or email already exist
	existingUser, _ := UserByName(user.Name)
	if existingUser.Name != "" {
		return errors.New("Name already used")
	}

	existingUser, _ = UserByEmail(user.Email)
	if existingUser.Email != "" {
		return errors.New("Email already used")
	}

	statement := "INSERT INTO users (uuid, name, email, password, created_at, discord_id, google_id, facebook_id, github_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		errors.New("Error preparing statement: " + err.Error())
		return
	}
	defer stmt.Close()

	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		errors.New("Error hashing password: " + err.Error())
		return
	}

	uuid := createUUID()
	user.Uuid = uuid
	_, err = stmt.Exec(uuid, user.Name, user.Email, bcryptPassword, time.Now(), user.DiscordID, user.GoogleID, user.FacebookID, user.GithubID)
	return
}

func UserByEmail(email string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("SELECT uuid, name, email, password, created_at, isAdmin, discord_id, google_id, facebook_id, github_id FROM users WHERE email = ?", email).
		Scan(&user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.IsAdmin, &user.DiscordID, &user.GoogleID, &user.FacebookID, &user.GithubID)
	return user, err
}

func UserByName(name string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("SELECT uuid, name, email, password, created_at, isAdmin, discord_id, google_id, facebook_id, github_id FROM users WHERE name = ?", name).
		Scan(&user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.IsAdmin, &user.DiscordID, &user.GoogleID, &user.FacebookID, &user.GithubID)
	return user, err
}

func UserByUUID(uuid string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("SELECT uuid, name, email, password, created_at, isAdmin, discord_id, google_id, facebook_id, github_id FROM users WHERE uuid = ?", uuid).
		Scan(&user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.IsAdmin, &user.DiscordID, &user.GoogleID, &user.FacebookID, &user.GithubID)
	return user, err
}

func UserByDiscordID(discordID string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("SELECT uuid, name, email, password, created_at, isAdmin, discord_id, google_id, facebook_id, github_id FROM users WHERE discord_id = ?", discordID).
		Scan(&user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.IsAdmin, &user.DiscordID, &user.GoogleID, &user.FacebookID, &user.GithubID)
	return user, err
}

func UserByGoogleID(googleID string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("SELECT uuid, name, email, password, created_at, isAdmin, discord_id, google_id, facebook_id, github_id FROM users WHERE google_id = ?", googleID).
		Scan(&user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.IsAdmin, &user.DiscordID, &user.GoogleID, &user.FacebookID, &user.GithubID)
	return user, err
}

func UserByFacebookID(facebookID string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("SELECT uuid, name, email, password, created_at, isAdmin, discord_id, google_id, facebook_id, github_id FROM users WHERE facebook_id = ?", facebookID).
		Scan(&user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.IsAdmin, &user.DiscordID, &user.GoogleID, &user.FacebookID, &user.GithubID)
	return user, err
}

func UserByGithubID(githubID string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("SELECT uuid, name, email, password, created_at, isAdmin, discord_id, google_id, facebook_id, github_id FROM users WHERE github_id = ?", githubID).
		Scan(&user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.IsAdmin, &user.DiscordID, &user.GoogleID, &user.FacebookID, &user.GithubID)
	return user, err
}

func GetAllUsers() (users []User, err error) {
	rows, err := Db.Query("SELECT uuid, name, email, password, created_at, isAdmin, discord_id, google_id, facebook_id, github_id FROM users")
	if err != nil {
		return
	}
	for rows.Next() {
		user := User{}
		if err = rows.Scan(&user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.IsAdmin, &user.DiscordID, &user.GoogleID, &user.FacebookID, &user.GithubID); err != nil {
			return
		}
		users = append(users, user)
	}
	rows.Close()
	return
}

func (user *User) SetName(name string) (err error) {
	user.Name = name
	return
}

func (user *User) SetEmail(email string) (err error) {
	user.Email = email
	return
}

func (user *User) Update() (err error) {
	statement := "UPDATE users SET name = ?, email = ? WHERE Uuid = ?"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Email, user.Uuid)
	return err
}

func (user *User) Update2() (err error) {
	statement := "UPDATE users SET isAdmin = ? WHERE Uuid = ?"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.IsAdmin, user.Uuid)
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
		fmt.Println("Error preparing statement: ", err)
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

func (user *User) CheckPassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) == nil
}
 
func (user *User) CreateThread(topicUuId string, title string, desc string) (thread Thread, err error) {
	statement := "INSERT INTO threads (uuid, topic_uuid, user_uuid, title, desc, created_at) VALUES (?, ?, ?, ?, ?, ?)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	uuid := createUUID()
	_, err = stmt.Exec(uuid, topicUuId, user.Uuid, title, desc, time.Now())
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

func (user *User) CreateTopic(name, desc, categoryuuid string) (topic Topic, err error) {
	statement := "INSERT INTO topics (uuid, name, desc, category_uuid, created_at) VALUES (?, ?, ?, ?, ?)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	uuid := createUUID()
	_, err = stmt.Exec(uuid, name, desc, categoryuuid, time.Now())
	if err != nil {
		return
	}

	stmt.QueryRow(uuid).Scan(&topic.Uuid, &topic.Name, &topic.Description, &topic.CategoryUuId, &topic.CreatedAt)
	return
}


func (user *User) CreateComment(thread *Thread, body string) (*Comment, error) {
	statement := "INSERT INTO comments (uuid, body, user_uuid, thread_uuid, created_at) VALUES (?, ?, ?, ?, ?)"
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

	comment := &Comment{}
	stmt.QueryRow(uuid).Scan(&comment.Uuid, &comment.Body, &comment.UserUuId, &comment.ThreadUuId, &comment.CreatedAt)
	return comment, nil
}

func (user *User) GetName() string {
	return user.Name
}

func (user *User) GetEmail() string {
	return user.Email
}

func (user *User) GetUuid() string {
	return user.Uuid
}

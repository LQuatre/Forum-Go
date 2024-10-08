package models

import (
	"time"
)

type Thread struct {
	Uuid      string
	Title     string
	Desc 	  string
	TopicUuId string
	UserUuId  string
	CreatedAt time.Time	
	Comments  []Comment
}

// format the CreatedAt date to display nicely on the screen
func (thread *Thread) CreatedAtDate() string {
	return thread.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

// get the number of posts in a thread
func (thread *Thread) NumReplies() (count int) {
	rows, err := Db.Query("SELECT count(*) FROM posts WHERE thread_uuid = ?", thread.UserUuId)
	if err != nil {
		return
	}
	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			return
		}
	}
	rows.Close()
	return
}

// get posts to a thread
func (thread *Thread) Posts() (comments []Comment, err error) {
	rows, err := Db.Query("SELECT uuid, body, user_id, thread_id, created_at FROM posts WHERE thread_uuid = ?", thread.Uuid)
	if err != nil {
		return
	}
	for rows.Next() {
		comment := Comment{}
		if err = rows.Scan(&comment.Uuid, &comment.Body, &comment.UserUuId, &comment.ThreadUuId, &comment.CreatedAt); err != nil {
			return
		}
		comments = append(comments, comment)
	}
	rows.Close()
	return
}

// Get all threads in a topic
func ThreadsByTopicUUID(topicUuId string) (threads []Thread, err error) {
	rows, err := Db.Query("SELECT uuid, topic_uuid, user_uuid, title, desc, created_at FROM threads WHERE topic_uuid = ? ORDER BY created_at DESC", topicUuId)
	if err != nil {
		return
	}
	for rows.Next() {
		thread := Thread{}
		if err = rows.Scan(&thread.Uuid, &thread.TopicUuId, &thread.UserUuId, &thread.Title, &thread.Desc, &thread.CreatedAt); err != nil {
			return
		}
		threads = append(threads, thread)
	}
	rows.Close()
	return
}

// Get a thread by the UUID
func ThreadByUUID(uuid string) (thread Thread, err error) {
	thread = Thread{}
	err = Db.QueryRow("SELECT uuid, topic_uuid, user_uuid, title, desc, created_at FROM threads WHERE uuid = ?", uuid).
		Scan(&thread.Uuid, &thread.TopicUuId, &thread.UserUuId, &thread.Title, &thread.Desc, &thread.CreatedAt)
	return
}

// Get the user who started this thread
func (thread *Thread) User() (user User) {
	user = User{}
	Db.QueryRow("SELECT uuid, name, email, created_at FROM users WHERE uuid = ?", thread.UserUuId).
		Scan(&user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	return
}
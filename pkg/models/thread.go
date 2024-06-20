package models

import "time"

type Thread struct {
	Uuid      string
	Title     string
	TopicUuId int
	UserUuId  int
	CreatedAt time.Time
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
func (thread *Thread) Posts() (posts []Post, err error) {
	rows, err := Db.Query("SELECT uuid, body, user_id, thread_id, created_at FROM posts WHERE thread_uuid = ?", thread.Uuid)
	if err != nil {
		return
	}
	for rows.Next() {
		post := Post{}
		if err = rows.Scan(&post.Uuid, &post.Body, &post.UserUuId, &post.ThreadUuId, &post.CreatedAt); err != nil {
			return
		}
		posts = append(posts, post)
	}
	rows.Close()
	return
}

// Get all threads in a topic
func ThreadsByTopic(topicUuId int) (threads []Thread, err error) {
	rows, err := Db.Query("SELECT uuid, title, topic_uuid, user_id, created_at FROM threads WHERE topic_uuid = ? ORDER BY created_at DESC", topicUuId)
	if err != nil {
		return
	}
	for rows.Next() {
		thread := Thread{}
		if err = rows.Scan(&thread.Uuid, &thread.Title, &thread.TopicUuId, &thread.UserUuId, &thread.CreatedAt); err != nil {
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
	err = Db.QueryRow("SELECT uuid, title, topic_id, user_id, created_at FROM threads WHERE uuid = ?", uuid).
		Scan(&thread.Uuid, &thread.Title, &thread.TopicUuId, &thread.UserUuId, &thread.CreatedAt)
	return
}

// Get the user who started this thread
func (thread *Thread) User() (user User) {
	user = User{}
	Db.QueryRow("SELECT uuid, name, email, created_at FROM users WHERE uuid = ?", thread.UserUuId).
		Scan(&user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	return
}

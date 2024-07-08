package models

import "time"

type Comment struct {
	Uuid       string
	Body       string
	UserUuId   string
	ThreadUuId string
	CreatedAt  time.Time
	Author	   User
	Likes    int
	Dislikes int
}

func (post *Comment) CreatedAtDate() string {
	return post.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

// Get the user who wrote the post
func (post *Comment) User() (user User) {
	user = User{}
	Db.QueryRow("SELECT uuid, name, email, created_at FROM users WHERE uuid = ?", post.UserUuId).
		Scan(&user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	return
}

func GetCommentsByThreadUUID(threadUuId string) (comments []Comment, err error) {
	rows, err := Db.Query("SELECT uuid, body, user_uuid, thread_uuid, created_at FROM comments WHERE thread_uuid = ?", threadUuId)
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
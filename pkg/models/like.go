package models

type Like struct {
	UUID   string `json:"uuid"`
	UserID string `json:"user_id"`
	PostID string `json:"post_id"`
}

func (l *Like) Create() error {
	_, err := Db.Exec("INSERT INTO likes (uuid, user_id, post_id) VALUES ($1, $2, $3)", l.UUID, l.UserID, l.PostID)
	return err
}

func (l *Like) Delete() error {
	_, err := Db.Exec("DELETE FROM likes WHERE uuid=$1", l.UUID)
	return err
}

func GetLikeByPostID(postID string) (*Like, error) {
	row := Db.QueryRow("SELECT uuid, user_id, post_id FROM likes WHERE post_id=$1", postID)
	like := &Like{}
	err := row.Scan(&like.UUID, &like.UserID, &like.PostID)
	return like, err
}

func GetLikesByPostID(postID string) ([]*Like, error) {
	rows, err := Db.Query("SELECT uuid, user_id, post_id FROM likes WHERE post_id=$1", postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	likes := []*Like{}
	for rows.Next() {
		like := &Like{}
		if err := rows.Scan(&like.UUID, &like.UserID, &like.PostID); err != nil {
			return nil, err
		}
		likes = append(likes, like)
	}
	return likes, nil
}

func GetLikeByUserID(userID string) (*Like, error) {
	row := Db.QueryRow("SELECT uuid, user_id, post_id FROM likes WHERE user_id=$1", userID)
	like := &Like{}
	err := row.Scan(&like.UUID, &like.UserID, &like.PostID)
	return like, err
}

func GetLikeByUUID(uuid string) (*Like, error) {
	row := Db.QueryRow("SELECT uuid, user_id, post_id FROM likes WHERE uuid=$1", uuid)
	like := &Like{}
	err := row.Scan(&like.UUID, &like.UserID, &like.PostID)
	return like, err
}

func GetAllLikes() ([]*Like, error) {
	rows, err := Db.Query("SELECT uuid, user_id, post_id FROM likes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	likes := []*Like{}
	for rows.Next() {
		like := &Like{}
		if err := rows.Scan(&like.UUID, &like.UserID, &like.PostID); err != nil {
			return nil, err
		}
		likes = append(likes, like)
	}
	return likes, nil
}

func (l *Like) Update() error {
	_, err := Db.Exec("UPDATE likes SET user_id=$1, post_id=$2 WHERE uuid=$3", l.UserID, l.PostID, l.UUID)
	return err
}

// dislike

type Dislike struct {
	UUID   string `json:"uuid"`
	UserID string `json:"user_id"`
	PostID string `json:"post_id"`
}

func (d *Dislike) Create() error {
	_, err := Db.Exec("INSERT INTO dislikes (uuid, user_id, post_id) VALUES ($1, $2, $3)", d.UUID, d.UserID, d.PostID)
	return err
}

func (d *Dislike) Delete() error {
	_, err := Db.Exec("DELETE FROM dislikes WHERE uuid=$1", d.UUID)
	return err
}

func GetDislikeByPostID(postID string) (*Dislike, error) {
	row := Db.QueryRow("SELECT uuid, user_id, post_id FROM dislikes WHERE post_id=$1", postID)
	dislike := &Dislike{}
	err := row.Scan(&dislike.UUID, &dislike.UserID, &dislike.PostID)
	return dislike, err
}

func GetDislikesByPostID(postID string) ([]*Dislike, error) {
	rows, err := Db.Query("SELECT uuid, user_id, post_id FROM dislikes WHERE post_id=$1", postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	dislikes := []*Dislike{}
	for rows.Next() {
		dislike := &Dislike{}
		if err := rows.Scan(&dislike.UUID, &dislike.UserID, &dislike.PostID); err != nil {
			return nil, err
		}
		dislikes = append(dislikes, dislike)
	}
	return dislikes, nil
}

func GetDislikeByUserID(userID string) (*Dislike, error) {
	row := Db.QueryRow("SELECT uuid, user_id, post_id FROM dislikes WHERE user_id=$1", userID)
	dislike := &Dislike{}
	err := row.Scan(&dislike.UUID, &dislike.UserID, &dislike.PostID)
	return dislike, err
}

func GetDislikeByUUID(uuid string) (*Dislike, error) {
	row := Db.QueryRow("SELECT uuid, user_id, post_id FROM dislikes WHERE uuid=$1", uuid)
	dislike := &Dislike{}
	err := row.Scan(&dislike.UUID, &dislike.UserID, &dislike.PostID)
	return dislike, err
}

func GetAllDislikes() ([]*Dislike, error) {
	rows, err := Db.Query("SELECT uuid, user_id, post_id FROM dislikes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	dislikes := []*Dislike{}
	for rows.Next() {
		dislike := &Dislike{}
		if err := rows.Scan(&dislike.UUID, &dislike.UserID, &dislike.PostID); err != nil {
			return nil, err
		}
		dislikes = append(dislikes, dislike)
	}
	return dislikes, nil
}

func (d *Dislike) Update() error {
	_, err := Db.Exec("UPDATE dislikes SET user_id=$1, post_id=$2 WHERE uuid=$3", d.UserID, d.PostID, d.UUID)
	return err
}
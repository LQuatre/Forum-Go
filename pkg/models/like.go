package models

type Like struct {
	UUID     string `json:"uuid"`
	UserUUID string `json:"user_uuid"`
	PostUUID string `json:"post_uuid"`
	Value    int    `json:"value"`
}

func (l *Like) Create() error {
	l.UUID = createUUID()
	_, err := Db.Exec("INSERT INTO likes (uuid, user_uuid, post_uuid) VALUES ($1, $2, $3)", l.UUID, l.UserUUID, l.PostUUID)
	return err
}

func (l *Like) Delete() error {
	_, err := Db.Exec("DELETE FROM likes WHERE uuid=$1", l.UUID)
	return err
}

func GetLikeByPostUUID(PostUUID string) (*Like, error) {
	row := Db.QueryRow("SELECT uuid, user_uuid, post_uuid FROM likes WHERE post_uuid=$1", PostUUID)
	like := &Like{}
	err := row.Scan(&like.UUID, &like.UserUUID, &like.PostUUID)
	return like, err
}

func GetLikesByPostUUID(PostUUID string) ([]*Like, error) {
	rows, err := Db.Query("SELECT uuid, user_uuid, post_uuid FROM likes WHERE post_uuid=$1", PostUUID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	likes := []*Like{}
	for rows.Next() {
		like := &Like{}
		if err := rows.Scan(&like.UUID, &like.UserUUID, &like.PostUUID); err != nil {
			return nil, err
		}
		likes = append(likes, like)
	}
	return likes, nil
}

func GetLikeByUserUUID(UserUUID string) (*Like, error) {
	row := Db.QueryRow("SELECT uuid, user_uuid, post_uuid FROM likes WHERE user_uuid=$1", UserUUID)
	like := &Like{}
	err := row.Scan(&like.UUID, &like.UserUUID, &like.PostUUID)
	return like, err
}

func GetLikeByUUID(uuid string) (*Like, error) {
	row := Db.QueryRow("SELECT uuid, user_uuid, post_uuid FROM likes WHERE uuid=$1", uuid)
	like := &Like{}
	err := row.Scan(&like.UUID, &like.UserUUID, &like.PostUUID)
	return like, err
}

func GetLikeByPostAndUserUUID(PostUUID string, UserUUID string) (*Like, error) {
	row := Db.QueryRow("SELECT uuid, user_uuid, post_uuid FROM likes WHERE post_uuid=$1 AND user_uuid=$2", PostUUID, UserUUID)
	like := &Like{}
	err := row.Scan(&like.UUID, &like.UserUUID, &like.PostUUID)
	return like, err
}

func GetAllLikes() ([]*Like, error) {
	rows, err := Db.Query("SELECT uuid, user_uuid, post_uuid FROM likes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	likes := []*Like{}
	for rows.Next() {
		like := &Like{}
		if err := rows.Scan(&like.UUID, &like.UserUUID, &like.PostUUID); err != nil {
			return nil, err
		}
		likes = append(likes, like)
	}
	return likes, nil
}

func (l *Like) Update() error {
	_, err := Db.Exec("UPDATE likes SET user_uuid=$1, post_uuid=$2 WHERE uuid=$3", l.UserUUID, l.PostUUID, l.UUID)
	return err
}

// dislike

type Dislike struct {
	UUID     string `json:"uuid"`
	UserUUID string `json:"user_uuid"`
	PostUUID string `json:"post_uuid"`
	Value    int    `json:"value"`
}

func (d *Dislike) Create() error {
	_, err := Db.Exec("INSERT INTO dislikes (uuid, user_uuid, post_uuid) VALUES ($1, $2, $3)", d.UUID, d.UserUUID, d.PostUUID)
	return err
}

func (d *Dislike) Delete() error {
	_, err := Db.Exec("DELETE FROM dislikes WHERE uuid=$1", d.UUID)
	return err
}

func GetDislikeByPostUUID(PostUUID string) (*Dislike, error) {
	row := Db.QueryRow("SELECT uuid, user_uuid, post_uuid FROM dislikes WHERE post_uuid=$1", PostUUID)
	dislike := &Dislike{}
	err := row.Scan(&dislike.UUID, &dislike.UserUUID, &dislike.PostUUID)
	return dislike, err
}

func GetDislikesByPostUUID(PostUUID string) ([]*Dislike, error) {
	rows, err := Db.Query("SELECT uuid, user_uuid, post_uuid FROM dislikes WHERE post_uuid=$1", PostUUID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	dislikes := []*Dislike{}
	for rows.Next() {
		dislike := &Dislike{}
		if err := rows.Scan(&dislike.UUID, &dislike.UserUUID, &dislike.PostUUID); err != nil {
			return nil, err
		}
		dislikes = append(dislikes, dislike)
	}
	return dislikes, nil
}

func GetDislikeByUserUUID(UserUUID string) (*Dislike, error) {
	row := Db.QueryRow("SELECT uuid, user_uuid, post_uuid FROM dislikes WHERE user_uuid=$1", UserUUID)
	dislike := &Dislike{}
	err := row.Scan(&dislike.UUID, &dislike.UserUUID, &dislike.PostUUID)
	return dislike, err
}

func GetDislikeByUUID(uuid string) (*Dislike, error) {
	row := Db.QueryRow("SELECT uuid, user_uuid, post_uuid FROM dislikes WHERE uuid=$1", uuid)
	dislike := &Dislike{}
	err := row.Scan(&dislike.UUID, &dislike.UserUUID, &dislike.PostUUID)
	return dislike, err
}

func GetDislikeByPostAndUserUUID(PostUUID string, UserUUID string) (*Dislike, error) {
	row := Db.QueryRow("SELECT uuid, user_uuid, post_uuid FROM dislikes WHERE post_uuid=$1 AND user_uuid=$2", PostUUID, UserUUID)
	dislike := &Dislike{}
	err := row.Scan(&dislike.UUID, &dislike.UserUUID, &dislike.PostUUID)
	return dislike, err
}

func GetAllDislikes() ([]*Dislike, error) {
	rows, err := Db.Query("SELECT uuid, user_uuid, post_uuid FROM dislikes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	dislikes := []*Dislike{}
	for rows.Next() {
		dislike := &Dislike{}
		if err := rows.Scan(&dislike.UUID, &dislike.UserUUID, &dislike.PostUUID); err != nil {
			return nil, err
		}
		dislikes = append(dislikes, dislike)
	}
	return dislikes, nil
}

func (d *Dislike) Update() error {
	_, err := Db.Exec("UPDATE dislikes SET user_uuid=$1, post_uuid=$2 WHERE uuid=$3", d.UserUUID, d.PostUUID, d.UUID)
	return err
}
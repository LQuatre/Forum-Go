package models

import "time"

type Topic struct {
	Id         int
	Uuid       string
	Name       string
	CategoryId int
	CreatedAt  time.Time
}

// Create a new topic in a category
func CreateTopic(name string, categoryId int) (topic Topic, err error) {
	statement := "INSERT INTO topics (uuid, name, category_id, created_at) VALUES (?, ?, ?, ?)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	uuid := createUUID()
	_, err = stmt.Exec(uuid, name, categoryId, time.Now())
	if err != nil {
		return
	}

	topic = Topic{
		Uuid:       uuid,
		Name:       name,
		CategoryId: categoryId,
		CreatedAt:  time.Now(),
	}
	return
}

// Get all topics in a category
func TopicsByCategory(categoryId int) (topics []Topic, err error) {
	rows, err := Db.Query("SELECT id, uuid, name, category_id, created_at FROM topics WHERE category_id = ?", categoryId)
	if err != nil {
		return
	}
	for rows.Next() {
		topic := Topic{}
		if err = rows.Scan(&topic.Id, &topic.Uuid, &topic.Name, &topic.CategoryId, &topic.CreatedAt); err != nil {
			return
		}
		topics = append(topics, topic)
	}
	rows.Close()
	return
}

// Get a topic by UUID
func TopicByUUID(uuid string) (topic Topic, err error) {
	topic = Topic{}
	err = Db.QueryRow("SELECT id, uuid, name, category_id, created_at FROM topics WHERE uuid = ?", uuid).
		Scan(&topic.Id, &topic.Uuid, &topic.Name, &topic.CategoryId, &topic.CreatedAt)
	return
}

func GetAllTopics() (topics []Topic, err error) {
	rows, err := Db.Query("SELECT id, uuid, name, category_id, created_at FROM topics")
	if err != nil {
		return
	}
	for rows.Next() {
		topic := Topic{}
		if err = rows.Scan(&topic.Id, &topic.Uuid, &topic.Name, &topic.CategoryId, &topic.CreatedAt); err != nil {
			return
		}
		topics = append(topics, topic)
	}
	rows.Close()
	return
}
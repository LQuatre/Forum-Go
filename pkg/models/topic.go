package models

import (
	"fmt"
	"time"
)

type Topic struct {
	Uuid         string
	Name         string
	CategoryUuId int
	CreatedAt    time.Time
}

// Create a new topic in a category
func CreateTopic(name string, categoryUuId int) (topic Topic, err error) {
	statement := "INSERT INTO topics (uuid, name, category_uuid, created_at) VALUES (?, ?, ?, ?)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	uuid := createUUID()
	_, err = stmt.Exec(uuid, name, categoryUuId, time.Now())
	if err != nil {
		return
	}

	topic = Topic{
		Uuid:         uuid,
		Name:         name,
		CategoryUuId: categoryUuId,
		CreatedAt:    time.Now(),
	}
	return
}

// Get all topics in a category
func TopicsByCategory(categoryUuId int) (topics []Topic, err error) {
	rows, err := Db.Query("SELECT uuid, name, category_id, created_at FROM topics WHERE category_uuid = ?", categoryUuId)
	if err != nil {
		return
	}
	for rows.Next() {
		topic := Topic{}
		if err = rows.Scan(&topic.Uuid, &topic.Name, &topic.CategoryUuId, &topic.CreatedAt); err != nil {
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
	err = Db.QueryRow("SELECT uuid, name, category_uuid, created_at FROM topics WHERE uuid = ?", uuid).
		Scan(&topic.Uuid, &topic.Name, &topic.CategoryUuId, &topic.CreatedAt)
	return
}

func GetAllTopics() (topics []Topic, err error) {
	rows, err := Db.Query("SELECT uuid, name, category_uuid, created_at FROM topics")
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	for rows.Next() {
		topic := Topic{}
		if err = rows.Scan(&topic.Uuid, &topic.Name, &topic.CategoryUuId, &topic.CreatedAt); err != nil {
			fmt.Printf("Error: %v", err)
			return
		}
		topics = append(topics, topic)
	}
	rows.Close()
	fmt.Printf("Topics: %v", topics)
	return
}

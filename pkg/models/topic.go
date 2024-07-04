package models

import (
	"fmt"
	"time"
)

type Topic struct {
	Uuid         string
	Name         string
	Description  string
	CategoryUuId string
	CreatedAt    time.Time
	Threads      []Thread
}

func Topics() (topics []Topic, err error) {
	rows, err := Db.Query("SELECT uuid, name, category_uuid, created_at FROM topics")
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

// Create a new topic in a category
func CreateTopic(name string, categoryUuId string) (topic Topic, err error) {
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
	return
}

func DeleteTopic(uuid string) (err error) {
	statement := "DELETE FROM topics WHERE uuid = ?"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		fmt.Println("Error in models/topic.go: DeleteTopic()")
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	return
}

func (topic *Topic) Delete() (err error) {
	statement := "DELETE FROM topics WHERE uuid = ?"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		fmt.Println("Error in models/topic.go: Delete()")
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(topic.Uuid)
	return
}

func (topic *Topic) Update() (err error) {
	statement := "UPDATE topics SET name = ? WHERE uuid = ?"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		fmt.Println("Error in models/topic.go: Update()")
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(topic.Name, topic.Uuid)
	return
}

func TopicsFromCategoryUUID(uuid string) (topics []Topic, err error) {
	rows, err := Db.Query("SELECT uuid, name, category_uuid, created_at FROM topics WHERE category_uuid = ?", uuid)
	if err != nil {
		fmt.Println("Error in models/topic.go: TopicsFromCategoryUUID()")
		return
	}
	for rows.Next() {
		topic := Topic{}
		if err = rows.Scan(&topic.Uuid, &topic.Name, &topic.CategoryUuId, &topic.CreatedAt); err != nil {
			fmt.Println("Error in models/topic.go: TopicsFromCategoryUUID()")
			return
		}
		topics = append(topics, topic)
	}
	rows.Close()
	return
}

func GetTopicByUUID(uuid string) (topic Topic, err error) {
	topic = Topic{}
	err = Db.QueryRow("SELECT uuid, name, category_uuid, created_at FROM topics WHERE uuid = ?", uuid).
		Scan(&topic.Uuid, &topic.Name, &topic.CategoryUuId, &topic.CreatedAt)
	return
}
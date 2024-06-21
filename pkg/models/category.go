package models

import "time"

type Category struct {
	Uuid      string
	Name      string
	CreatedAt time.Time
}

func (category *Category) Create(name string) (err error) {
	statement := "INSERT INTO categories (uuid, name, created_at) VALUES (?, ?, ?)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	category.Uuid = createUUID()
	category.Name = name
	category.CreatedAt = time.Now()
	_, err = stmt.Exec(category.Uuid, category.Name, category.CreatedAt)
	return
}

// Get all categories
func Categories() (categories []Category, err error) {
	rows, err := Db.Query("SELECT uuid, name, created_at FROM categories")
	if err != nil {
		return
	}
	for rows.Next() {
		category := Category{}
		if err = rows.Scan(&category.Uuid, &category.Name, &category.CreatedAt); err != nil {
			return
		}
		categories = append(categories, category)
	}
	rows.Close()
	return
}

// Get a category by UUID
func CategoryByUUID(uuid string) (category Category, err error) {
	category = Category{}
	err = Db.QueryRow("SELECT uuid, name, created_at FROM categories WHERE uuid = ?", uuid).
		Scan(&category.Uuid, &category.Name, &category.CreatedAt)
	return
}

func GetAllCategories() (categories []Category, err error) {
	rows, err := Db.Query("SELECT uuid, name, created_at FROM categories")
	if err != nil {
		return
	}
	for rows.Next() {
		category := Category{}
		if err = rows.Scan(&category.Uuid, &category.Name, &category.CreatedAt); err != nil {
			return
		}
		categories = append(categories, category)
	}
	rows.Close()
	return
}

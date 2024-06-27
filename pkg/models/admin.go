package models

type AdminInfo struct {
	Users      []User
	Sessions   []Session
	Categories []Category
	Topics     []Topic
}

func Admin() (admin AdminInfo, err error) {
	users, err := GetAllUsers()
	if err != nil {
		return
	}
	sessions, err := GetAllSessions()
	if err != nil {
		return
	}
	categories, err := GetAllCategories()
	if err != nil {
		return
	}
	topics, err := GetAllTopics()
	if err != nil {
		return
	}
	admin = AdminInfo{Users: users, Sessions: sessions, Categories: categories, Topics: topics}
	return
}

func AdminUser() (admin User, err error) {
	admin = User{}
	return
}
package models

type User struct {
	ID         uint
	Username   string
	Password   string
	CreateTime int64
}

func (User) TableName() string {
	return "go_users"
}

package models

type UserModel struct {
	ID         uint
	Username   string
	Password   string
	CreateTime int64
}

func (UserModel) TableName() string {
	return "go_users"
}

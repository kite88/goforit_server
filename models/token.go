package models

type TokenModel struct {
	ID     uint
	Token  string
	Expire int64
	UserId uint
}

func (TokenModel) TableName() string {
	return "go_tokens"
}

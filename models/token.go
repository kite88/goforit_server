package models

type TokenModel struct {
	ID     uint64
	Token  string
	Expire int64
	UserId uint64
}

func (TokenModel) TableName() string {
	return "go_tokens"
}

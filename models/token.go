package models

type Token struct {
	ID     uint
	Token  string
	Expire int64
	UserId uint
}

func (Token) TableName() string {
	return "go_tokens"
}

package models

type MemoClassifyModel struct {
	ID         uint `json:"id"`
	Name       string `json:"name"`
	CreateTime int64 `json:"create_time"`
	UserId     uint `json:"user_id"`
	DeleteTime int64 `json:"delete_time"`
}

func (MemoClassifyModel) TableName() string {
	return "go_memo_classify"
}

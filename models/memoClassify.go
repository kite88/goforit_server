package models

type MemoClassifyModel struct {
	ID         uint64 `json:"id"`
	Name       string `json:"name"`
	CreateTime int64  `json:"create_time"`
	UserId     uint64 `json:"user_id"`
	DeleteTime int64  `json:"-"`
}

func (MemoClassifyModel) TableName() string {
	return "go_memo_classify"
}

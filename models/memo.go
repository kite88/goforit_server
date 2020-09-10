package models

type MemosModel struct {
	ID         uint64 `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
	UserId     uint64 `json:"user_id"`
	ClassifyId uint64 `json:"classify_id"`
	UpdateTime int64  `json:"update_time"`
	DeleteTime int64  `json:"delete_time"`
}

func (MemosModel) TableName() string {
	return "go_memos"
}

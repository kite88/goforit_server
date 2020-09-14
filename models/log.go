package models

type LogModel struct {
	ID         uint64 `json:"id"`
	LogId      string `json:"log_id"`
	CreateTime int64  `json:"create_time"`
	UserId     uint64 `json:"user_id"`
	TypeId     uint64 `json:"type_id"`
	LogText    string `json:"log_text"`
}

func (LogModel) TableName() string {
	return "go_logs"
}

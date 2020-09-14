package models

type LogTypeModel struct {
	ID uint64 `json:"id"`
	Name string `json:"name"`
}

func (LogTypeModel) TableName() string {
	return "go_log_type"
}

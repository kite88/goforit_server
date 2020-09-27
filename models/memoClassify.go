package models

type MemoClassifyModel struct {
	ID              uint64               `json:"id"`
	Name            string               `json:"name"`
	CreateTime      int64                `json:"create_time"`
	UserId          uint64               `json:"user_id"`
	DeleteTime      int64                `json:"-"`
	Pid             uint64               `json:"pid"`
	PidMemoClassify PidMemoClassifyModel `json:"pid_memo_classify" gorm:"FOREIGNKEY:ID;ASSOCIATION_FOREIGNKEY:Pid"`
}

type PidMemoClassifyModel struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

func (MemoClassifyModel) TableName() string {
	return "go_memo_classify"
}

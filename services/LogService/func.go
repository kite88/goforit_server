package LogService

import (
	"goforit/db"
	"goforit/models"
	"goforit/utils"
	"time"
)

func CreateLogs(userId uint64, typeId uint64, logText string) bool {
	err := db.DB().Create(&models.LogModel{
		LogId:      utils.GenXid(),
		CreateTime: time.Now().Unix(),
		UserId:     userId,
		TypeId:     typeId,
		LogText:    "" + logText,
	}).Error
	if err == nil {
		return true
	}
	return false
}

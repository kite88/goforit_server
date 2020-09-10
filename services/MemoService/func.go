package MemoService

import (
	"goforit/db"
	"goforit/models"
	"time"
)

var (
	commonWhere string = "user_id = ? AND delete_time = ?"
)

func GetMemo(userId uint) []models.MemosModel {
	var list []models.MemosModel
	db.DB().Where(commonWhere, userId, 0).Find(&list)
	return list
}

func CreateMemo(title string, content string, classifyId uint, userId uint) bool {
	err := db.DB().Create(&models.MemosModel{
		Title:      title,
		Content:    content,
		CreateTime: time.Now().Unix(),
		UserId:     userId,
		ClassifyId: classifyId,
	}).Error
	if nil == err {
		return true
	}
	return false
}

func EditMemo(id uint, title string, content string, classifyId uint, userId uint) int64 {
	var count int64
	count = db.DB().Model(&models.MemosModel{}).Where(commonWhere, userId, 0).Where("id = ?", id).Updates(&models.MemosModel{
		Title:      title,
		Content:    content,
		ClassifyId: classifyId,
		UpdateTime: time.Now().Unix(),
	}).RowsAffected
	return count
}

func DeleteMemo(id uint, userId uint) int64 {
	var count int64
	count = db.DB().Model(&models.MemosModel{}).
		Where(commonWhere, userId, 0).Where("id = ?", id).
		Update("delete_time", time.Now().Unix()).RowsAffected
	return count
}

func FindMemoClassify(id uint, userId uint) (models.MemoClassifyModel, int) {
	var memoClassify models.MemoClassifyModel
	var count int
	db.DB().Where(commonWhere, userId, 0).Where("id = ?", id).First(&memoClassify).Count(&count)
	return memoClassify, count
}

func GetMemoClassify(userId uint) []models.MemoClassifyModel {
	var list []models.MemoClassifyModel
	db.DB().Where(commonWhere, userId, 0).Find(&list)
	return list
}

func CreateMemoClassify(name string, userId uint) bool {
	err := db.DB().Create(&models.MemoClassifyModel{
		Name:       name,
		CreateTime: time.Now().Unix(),
		UserId:     userId,
	}).Error
	if err == nil {
		return true
	}
	return false
}

func EditMemoClassify(name string, id uint, userId uint) int64 {
	var count int64
	count = db.DB().Model(&models.MemoClassifyModel{}).
		Where(commonWhere, userId, 0).Where("id = ?", id).
		Updates(models.MemoClassifyModel{
			Name: name,
		}).RowsAffected
	return count
}

func DeleteMemoClassify(id uint, userId uint) int64 {
	var count int64
	count = db.DB().Model(&models.MemoClassifyModel{}).
		Where(commonWhere, userId, 0).Where("id = ?", id).
		Update("delete_time", time.Now().Unix()).RowsAffected
	return count
}

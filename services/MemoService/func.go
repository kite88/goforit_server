package MemoService

import (
	"fmt"
	"goforit/db"
	"goforit/models"
	"goforit/utils/Page"
	"time"
)

var (
	commonWhere string = "user_id = ? AND delete_time = ?"
)

func GetMemo(userId uint64, title string, classifyId uint64, pageIndex uint64, pageSize uint64) (result Page.Lists) {
	var MemosModel models.MemosModel
	var list []models.MemosModel
	var count uint64
	var whereTitle = ""
	var whereClassifyId = ""
	if title != "" {
		whereTitle = "title LIKE " + "'%" + title + "%'"
	}
	if classifyId > 0 {
		whereClassifyId = fmt.Sprintf("classify_id = %d", classifyId)
	}
	db.DB().Model(&MemosModel).Where(commonWhere, userId, 0).
		Where(whereTitle).Where(whereClassifyId).
		Count(&count)
	offset, length, pageResult := Page.Make(count, pageIndex, pageSize)
	db.DB().Model(&MemosModel).
		Where(commonWhere, userId, 0).
		Where(whereTitle).Where(whereClassifyId).
		Offset(offset).Limit(length).
		Preload("MemoClassify").
		Order("create_time desc").
		Find(&list)
	result.Page = pageResult
	result.List = list
	return result
}

func CreateMemo(title string, content string, classifyId uint64, userId uint64) bool {
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

func EditMemo(id uint64, title string, content string, classifyId uint64, userId uint64) int64 {
	var count int64
	count = db.DB().Model(&models.MemosModel{}).Where(commonWhere, userId, 0).Where("id = ?", id).Updates(&models.MemosModel{
		Title:      title,
		Content:    content,
		ClassifyId: classifyId,
		UpdateTime: time.Now().Unix(),
	}).RowsAffected
	return count
}

func DeleteMemo(id uint64, userId uint64) int64 {
	var count int64
	count = db.DB().Model(&models.MemosModel{}).
		Where(commonWhere, userId, 0).Where("id = ?", id).
		Update("delete_time", time.Now().Unix()).RowsAffected
	return count
}

func FindMemoClassify(id uint64, userId uint64, isPid bool) (models.MemoClassifyModel, int64) {
	var memoClassify models.MemoClassifyModel
	var count int64
	var pidWhere string
	if isPid == true {
		pidWhere = "pid = 0"
	}
	db.DB().Where(commonWhere, userId, 0).
		Where("id = ?", id).
		Where(pidWhere).
		First(&memoClassify).Count(&count)
	return memoClassify, count
}

func GetMemoClassify(userId uint64) []models.MemoClassifyModel {
	var list []models.MemoClassifyModel
	db.DB().Where(commonWhere, userId, 0).Find(&list)
	return recursionMemoClassify(0, list)
}

func recursionMemoClassify(pid uint64, list []models.MemoClassifyModel) []models.MemoClassifyModel {
	var nList []models.MemoClassifyModel
	for _, value := range list {
		if pid == value.Pid {
			child := recursionMemoClassify(value.ID, list)
			nList = append(append(nList, value), child...)
		}
	}
	return nList
}

func CreateMemoClassify(pid uint64, name string, userId uint64) bool {
	err := db.DB().Create(&models.MemoClassifyModel{
		Name:       name,
		CreateTime: time.Now().Unix(),
		UserId:     userId,
		Pid:        pid,
	}).Error
	if err == nil {
		return true
	}
	return false
}

func EditMemoClassify(name string, id uint64, userId uint64) int64 {
	var count int64
	count = db.DB().Model(&models.MemoClassifyModel{}).
		Where(commonWhere, userId, 0).Where("id = ?", id).
		Updates(models.MemoClassifyModel{
			Name: name,
		}).RowsAffected
	return count
}

func DeleteMemoClassify(id uint64, userId uint64) int64 {
	var count int64
	count = db.DB().Model(&models.MemoClassifyModel{}).
		Where(commonWhere, userId, 0).Where("id = ?", id).
		Update("delete_time", time.Now().Unix()).RowsAffected
	return count
}

func FindMemoClassifyChild(pid uint64, userId uint64) ([]models.MemoClassifyModel, int64) {
	var list []models.MemoClassifyModel
	var count int64
	db.DB().Where(commonWhere, userId, 0).
		Where("pid = ?", pid).
		Find(&list).Count(&count)
	return list, count
}

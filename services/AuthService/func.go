package AuthService

import (
	"goforit/db"
	"goforit/models"
	"goforit/utils"
	"time"
)

func Logout(userId uint) {
	var TokenModel models.Token
	db.DB().Model(&TokenModel).Where("user_id = ?", userId).Update("expire", time.Now().Unix())
}

func Login(username string, password string) (models.User, int) {
	var user models.User
	err := db.DB().Where("username = ?", username).First(&user).Error
	if err != nil {
		return user, 0
	}
	if utils.Md5(password) == user.Password && user.Username == username {
		return user, 1
	}
	return user, 0
}

func Token(userId uint, token string) string {
	var TokenModel models.Token
	var count int
	db.DB().Model(&TokenModel).Where("user_id = ?", userId).Count(&count)
	if count == 0 {
		err := db.DB().Create(&models.Token{
			Token:  token,
			Expire: time.Now().Unix() + 24*3600,
			UserId: userId,
		}).Error
		if err == nil {
			return token
		}
	} else {
		err := db.DB().Model(&TokenModel).Where("user_id = ?", userId).Updates(models.Token{
			Token:  token,
			Expire: time.Now().Unix() + 24*3600,
		}).Error
		if err == nil {
			return token
		}
	}
	return ""
}

func GetUserIdByToken(token string) uint {
	var TokenModel models.Token
	err := db.DB().Where("token = ?", token).First(&TokenModel).Error
	if err == nil {
		if TokenModel.Expire < time.Now().Unix() {
			return 0
		}
		return TokenModel.UserId
	}
	return 0
}
package db

import (
	"github.com/qingxunying/douyin/constdef"
	"time"
)

type UserInfo struct {
	UserId     int64     `gorm:"user_id"`
	UserName   string    `gorm:"user_name"`
	Password   string    `gorm:"password"`
	Status     int       `gorm:"column:status;default:null"`
	CreateTime time.Time `gorm:"column:create_time;default:null"`
	UpdateTime time.Time `gorm:"column:update_time;default:null"`
}

func AddUserInfo(userId int64, userName, password string) *UserInfo {
	userInfo := UserInfo{
		UserId:   userId,
		UserName: userName,
		Password: password,
	}
	db := GetDB().Table(constdef.UserInfoTable)
	db.Create(&userInfo)
	return &userInfo
}

func GetUserInfoByUserName(userName string) *UserInfo {
	userInfo := &UserInfo{}
	db := GetDB().Table(constdef.UserInfoTable)
	db.Where("user_name = ? and status = ?", userName, constdef.NameOn).First(userInfo)
	if userInfo.UserId == 0 {
		return nil
	}
	return userInfo
}

func GetUserInfoByUserId(userId int64) *UserInfo {
	userInfo := &UserInfo{}
	db := GetDB().Table(constdef.UserInfoTable)
	db.Where("user_id = ? and status = ?", userId, constdef.NameOn).First(userInfo)
	if userInfo.UserId == 0 {
		return nil
	}
	return userInfo
}

func GetUserInfoByPassword(userName string, password string) *UserInfo {
	userInfo := &UserInfo{}
	db := GetDB().Table(constdef.UserInfoTable)
	db.Where("user_name = ? and password = ? and status = ?", userName, password, constdef.NameOn).First(userInfo)
	if userInfo.UserId == 0 {
		return nil
	}
	return userInfo
}

func GetUserNameByUserId(userId int64) string {
	userInfo := UserInfo{}
	db := GetDB().Table(constdef.UserInfoTable)
	db.Where("user_id = ? and status = ?", userId, constdef.NameOn).First(&userInfo)
	return userInfo.UserName
}

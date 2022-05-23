package db

import (
	"github.com/qingxunying/douyin/constdef"
	"time"
)

type FollowInfo struct {
	UserId       int64     `gorm:"user_id"`
	FollowUserId int64     `gorm:"follow_user_id"`
	Status       int       `gorm:"status"`
	CreateTime   time.Time `gorm:"column:create_time;default:null"`
	UpdateTime   time.Time `gorm:"column:update_time;default:null"`
}

func GetFollowCount(userId int64) int64 {
	var count int64
	db.Table(constdef.FollowInfoTable).Where("user_id = ? and status = ?", userId, constdef.FollowOn).Count(&count)
	return count
}

func GetFollowerCount(followUserId int64) int64 {
	var count int64
	db.Table(constdef.FollowInfoTable).Where("follow_user_id = ? and status = ?", followUserId, constdef.FollowOn).Count(&count)
	return count
}

func IsFollowedRelation(userId int64, anotherUserId int64) bool {
	var count int64
	db.Table(constdef.FollowInfoTable).Where("user_id = ? and follow_user_id = ? and status = ?", anotherUserId, userId, constdef.FollowOn).Count(&count)
	return count >= 1
}

func GetFollowInfo(userId int64, followUserId int64) *FollowInfo {
	var followInfo FollowInfo
	db.Table(constdef.FollowInfoTable).
		Where("user_id = ? and follow_user_id = ?", userId, followUserId).
		First(&followInfo)
	if followInfo.UserId == 0 || followInfo.FollowUserId == 0 {
		return nil
	}
	return &followInfo
}

func AddFollowInfo(userId int64, followUserId int64, actionType int) {
	followInfo := &FollowInfo{
		UserId:       userId,
		FollowUserId: followUserId,
		Status:       actionType,
	}
	db.Table(constdef.FollowInfoTable).Create(followInfo)
}

func UpdateFollowInfo(userId int64, followUserId int64, actionType int) {
	db.Table(constdef.FollowInfoTable).
		Where("user_id = ? and follow_user_id = ?", userId, followUserId).
		Update(map[string]interface{}{
			"status": actionType,
		})
}

func GetAllFollowUser(userId int64) []FollowInfo {
	var followInfos []FollowInfo
	db.Table(constdef.FollowInfoTable).Where("user_id = ? and status = ?", userId, constdef.FollowOn).Find(&followInfos)
	return followInfos
}

func GetAllFollowerUser(userId int64) []FollowInfo {
	var followerInfos []FollowInfo
	db.Table(constdef.FollowInfoTable).Where("follow_user_id = ? and status = ?", userId, constdef.FollowOn).Find(&followerInfos)
	return followerInfos
}

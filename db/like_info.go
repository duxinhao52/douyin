package db

import (
	"github.com/qingxunying/douyin/constdef"
	"time"
)

type LikeInfo struct {
	UserId     int64     `gorm:"user_id"`
	VideoId    int64     `gorm:"video_id"`
	Status     int       `gorm:"status"`
	CreateTime time.Time `gorm:"column:create_time;default:null"`
	UpdateTime time.Time `gorm:"column:update_time;default:null"`
}

func GetLikeInfo(userId int64, videoId int64) *LikeInfo {
	var likeInfo LikeInfo
	db.Table(constdef.LikeInfoTable).
		Where("user_id = ? and video_id = ?", userId, videoId).
		First(&likeInfo)
	if likeInfo.UserId == 0 {
		return nil
	}
	return &likeInfo
}

func AddLikeInfo(userId int64, videoId int64, actionType int) {
	likeInfo := &LikeInfo{
		UserId:  userId,
		VideoId: videoId,
		Status:  actionType,
	}
	db.Table(constdef.LikeInfoTable).Create(likeInfo)
}

func UpdateLikeInfo(userId int64, videoId int64, actionType int) {
	db.Table(constdef.LikeInfoTable).
		Where("user_id = ? and video_id = ?", userId, videoId).
		Update(map[string]interface{}{
			"status": actionType,
		})
}

func GetLikeCount(videoId int64) int64 {
	var count int64
	db.Table(constdef.LikeInfoTable).Where("video_id = ? and status = ?", videoId, constdef.LikeOn).Count(&count)
	return count
}

func IsLikedRelation(userId int64, videoId int64) bool {
	var count int64
	db.Table(constdef.LikeInfoTable).Where("user_id = ? and video_id = ? and status = ?", userId, videoId, constdef.LikeOn).Count(&count)
	return count >= 1
}

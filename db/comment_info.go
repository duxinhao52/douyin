package db

import (
	"github.com/qingxunying/douyin/constdef"
	"time"
)

type CommentInfo struct {
	UserId     int64     `gorm:"user_id"`
	CommentId  int64     `gorm:"comment_id"`
	VideoId    int64     `gorm:"video_id"`
	Status     int       `gorm:"status"`
	Content    string    `gorm:"content"`
	CreateTime time.Time `gorm:"column:create_time;default:null"`
	UpdateTime time.Time `gorm:"column:update_time;default:null"`
}

func GetCommentCount(videoId int64) int64 {
	var count int64
	db.Table(constdef.CommentInfoTable).Where("video_id = ? and status = ?", videoId, constdef.CommentOn).Count(&count)
	return count
}

func AddComment(userId int64, videoId int64, commentId int64, content string) *CommentInfo {
	commentInfo := CommentInfo{
		UserId:    userId,
		CommentId: commentId,
		VideoId:   videoId,
		Status:    constdef.CommentOn,
		Content:   content,
	}
	db := GetDB().Table(constdef.CommentInfoTable)
	db.Create(&commentInfo)
	return &commentInfo
}

func DeleteComment(commentId int64) *CommentInfo {
	var commentInfo CommentInfo
	db := GetDB().Table(constdef.CommentInfoTable)
	db.Where("comment_id = ?", commentId).First(&commentInfo).Update(map[string]interface{}{
		"status": constdef.CommentOff,
	})
	if commentInfo.CommentId == 0 {
		return nil
	}
	return &commentInfo
}

func GetCommentInfo(videoId int64) []CommentInfo {
	var commentInfos []CommentInfo
	db.Table(constdef.CommentInfoTable).
		Where("video_id = ? and status = ?", videoId, constdef.PublishOn).
		Find(&commentInfos)
	return commentInfos
}

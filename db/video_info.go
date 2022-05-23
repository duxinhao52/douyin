package db

import (
	"github.com/qingxunying/douyin/constdef"
	"time"
)

type VideoInfo struct {
	VideoId    int64     `gorm:"video_id"`
	UserId     int64     `gorm:"user_id"`
	Title      string    `gorm:"title"`
	Url        string    `gorm:"url"`
	CoverUrl   string    `gorm:"column:cover_url;default:null"`
	Status     int       `gorm:"column:status;default:null"`
	CreateTime time.Time `gorm:"column:create_time;default:null"`
	UpdateTime time.Time `gorm:"column:update_time;default:null"`
}

func AddVideoInfo(videoId int64, userId int64, title, url string) *VideoInfo {
	videoInfo := &VideoInfo{
		VideoId: videoId,
		UserId:  userId,
		Title:   title,
		Url:     url,
	}
	db := GetDB().Table(constdef.VideoInfoTable)
	db.Create(videoInfo)
	return videoInfo
}

func GetAllVideoInfo(latestTime int64, currentTime int64) []VideoInfo {
	var videoInfos []VideoInfo
	db := GetDB().Table(constdef.VideoInfoTable)
	if latestTime == currentTime {
		db.Where("status = ? and create_time <= ?", constdef.PublishOn, time.Unix(latestTime, 0)).
			Order("create_time DESC").
			Limit(constdef.MaxVideoCount).
			Find(&videoInfos)
	} else {
		db.Where("status = ? and create_time < ?", constdef.PublishOn, time.Unix(latestTime, 0)).
			Order("create_time DESC").
			Limit(constdef.MaxVideoCount).
			Find(&videoInfos)
	}
	return videoInfos
}

func GetFavoriteVideoInfo(userId int64) []VideoInfo {
	var videoInfos []VideoInfo
	db.Raw("select video.video_id, video.user_id, video.title, video.url, video.cover_url, video.status, video.create_time, video.update_time from video_info video left join like_info liker on video.video_id = liker.video_id where liker.user_id = ? and liker.status = ? and video.status = ?", userId, constdef.LikeOn, constdef.PublishOn).Find(&videoInfos)
	return videoInfos
}

func GetPublishVideoInfo(userId int64) []VideoInfo {
	var videoInfos []VideoInfo
	db := GetDB().Table(constdef.VideoInfoTable)
	db.Where("user_id = ? and status = ?", userId, constdef.PublishOn).Find(&videoInfos)
	return videoInfos
}

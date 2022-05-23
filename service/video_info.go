package service

import (
	"github.com/qingxunying/douyin/db"
	"github.com/qingxunying/douyin/model"
	"github.com/sirupsen/logrus"
)

func GetAllVideoList(userId int64, latestTime int64, currentTime int64) ([]model.Video, int64) {
	var videoList []model.Video
	videoInfos := db.GetAllVideoInfo(latestTime, currentTime)
	for _, videoInfo := range videoInfos {
		videoList = append(videoList, GetVideo(videoInfo, userId, false))
	}
	nextTime := currentTime
	if len(videoInfos) != 0 {
		nextTime = videoInfos[len(videoInfos)-1].CreateTime.Unix()
	}
	return videoList, nextTime
}

func GetFavoriteVideoList(userId int64) []model.Video {
	var videoList []model.Video
	videoInfos := db.GetFavoriteVideoInfo(userId)
	logrus.Infof("[GetFavoriteVideoList] videoInfos=%+v", videoInfos)
	for _, videoInfo := range videoInfos {
		videoList = append(videoList, GetVideo(videoInfo, userId, true))
	}
	return videoList
}

func GetPublishVideoList(userId int64) []model.Video {
	var videoList []model.Video
	videoInfos := db.GetPublishVideoInfo(userId)
	for _, videoInfo := range videoInfos {
		videoList = append(videoList, GetVideo(videoInfo, userId, false))
	}
	return videoList
}

func GetVideo(videoInfo db.VideoInfo, userId int64, liked bool) model.Video {
	favoriteCount := db.GetLikeCount(userId)
	commentCount := db.GetCommentCount(userId)
	isFavorite := false
	if liked || userId != 0 && db.IsLikedRelation(userId, videoInfo.VideoId) {
		isFavorite = true
	}
	video := model.Video{
		Id:            videoInfo.VideoId,
		PlayUrl:       videoInfo.Url,
		Author:        *GetUser(videoInfo.UserId, userId),
		FavoriteCount: favoriteCount,
		CommentCount:  commentCount,
		CoverUrl:      videoInfo.CoverUrl,
		IsFavorite:    isFavorite,
	}
	return video
}

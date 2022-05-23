package service

import "github.com/qingxunying/douyin/db"

func AddLikeInfo(userId int64, videoId int64, actionType int) {
	likeInfo := db.GetLikeInfo(userId, videoId)
	if likeInfo == nil {
		db.AddLikeInfo(userId, videoId, actionType)
	} else if likeInfo.Status != actionType {
		db.UpdateLikeInfo(userId, videoId, actionType)
	}
	return
}

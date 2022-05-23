package service

import (
	"github.com/qingxunying/douyin/db"
)

func AddFollowInfo(userId, followUserId int64, actionType int) {
	followInfo := db.GetFollowInfo(userId, followUserId)
	if followInfo == nil {
		db.AddFollowInfo(userId, followUserId, actionType)
	} else if followInfo.Status != actionType {
		db.UpdateFollowInfo(userId, followUserId, actionType)
	}
	return
}

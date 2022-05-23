package service

import (
	"errors"
	"github.com/qingxunying/douyin/constdef"
	"github.com/qingxunying/douyin/db"
	"github.com/qingxunying/douyin/model"
	"github.com/qingxunying/douyin/util"
)

func Comment(actionType int, userId int64, videoId int64, commentId int64, content string) (*model.Comment, error) {
	if actionType == constdef.CommentOn {
		commentInfo := AddComment(userId, videoId, content)
		comment := &model.Comment{
			Id:         commentInfo.CommentId,
			User:       *GetUser(userId, userId),
			Content:    commentInfo.Content,
			CreateDate: commentInfo.CreateTime.Format("01-02"),
		}
		return comment, nil
	} else if actionType == constdef.CommentOff {
		commentInfo := DeleteComment(commentId)
		if commentInfo == nil {
			return nil, errors.New("delete error")
		} else {
			comment := &model.Comment{
				Id:         commentInfo.CommentId,
				User:       *GetUser(userId, userId),
				Content:    commentInfo.Content,
				CreateDate: commentInfo.CreateTime.Format("01-02"),
			}
			return comment, nil
		}
	}
	return nil, nil
}

func AddComment(userId int64, videoId int64, content string) *db.CommentInfo {
	commentId := util.CreateUuid()
	return db.AddComment(userId, videoId, commentId, content)
}

func DeleteComment(commentId int64) *db.CommentInfo {
	return db.DeleteComment(commentId)
}

func GetVideoComment(videoId int64, userId int64) []model.Comment {
	var comments []model.Comment
	commentInfos := db.GetCommentInfo(videoId)
	for _, commentInfo := range commentInfos {
		comments = append(comments, model.Comment{
			Id:         commentInfo.CommentId,
			User:       *GetUser(commentInfo.UserId, userId),
			Content:    commentInfo.Content,
			CreateDate: commentInfo.CreateTime.Format("01-02"),
		})
	}
	return comments
}

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/qingxunying/douyin/model"
	"github.com/qingxunying/douyin/service"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type CommentResponse struct {
	model.Response
	Comment model.Comment `json:"comment"`
}

type CommentListResponse struct {
	model.Response
	CommentList []model.Comment `json:"comment_list,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	userId, userName := service.ParseToken(token)
	if userId == 0 || userName == "" {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	actionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 64)
	content := c.Query("comment_text")
	commentId, _ := strconv.ParseInt(c.Query("comment_id"), 10, 64)
	logrus.Infof("[CommentAction] videoId=%v, actionType=%v, content=%v, commentId=%v", videoId, actionType, content, commentId)
	comments, err := service.Comment(int(actionType), userId, videoId, commentId, content)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	logrus.Infof("[CommentAction] comments=%+v", *comments)
	c.JSON(http.StatusOK, CommentResponse{
		Response: model.Response{StatusCode: 0},
		Comment:  *comments,
	})
	return
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	token := c.Query("token")
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	userId, userName := service.ParseToken(token)
	if userId == 0 || userName == "" {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	comments := service.GetVideoComment(videoId, userId)
	logrus.Infof("[CommentAction] comments=%v", comments)
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    model.Response{StatusCode: 0},
		CommentList: comments,
	})
	return
}

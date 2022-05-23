package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/qingxunying/douyin/model"
	"github.com/qingxunying/douyin/service"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	userId, userName := service.ParseToken(token)
	logrus.Infof("[FavoriteAction] userId=%v", userId)
	if userId == 0 || userName == "" {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	actionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 64)
	service.AddLikeInfo(userId, videoId, int(actionType))

	c.JSON(http.StatusOK, model.Response{StatusCode: 0})
	return
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	token := c.Query("token")
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if !service.CheckToken(userId, token) {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "check token failed",
		})
		return
	}

	logrus.Infof("[FavoriteList] userId=%v", userId)
	videoList := service.GetFavoriteVideoList(userId)
	logrus.Infof("[FavoriteList] videoList=%+v", videoList)
	c.JSON(http.StatusOK, VideoListResponse{
		Response: model.Response{
			StatusCode: 0,
		},
		VideoList: videoList,
	})
	return
}

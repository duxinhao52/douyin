package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/qingxunying/douyin/db"
	"github.com/qingxunying/douyin/model"
	"github.com/qingxunying/douyin/oss"
	"github.com/qingxunying/douyin/service"
	"github.com/qingxunying/douyin/util"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type VideoListResponse struct {
	model.Response
	VideoList []model.Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")
	title := c.PostForm("title")
	userId, userName := service.ParseToken(token)
	if userId == 0 || userName == "" {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	url, err := oss.UpVideoToOss(data, userId)
	logrus.Infof("[Publish] url=%v", url)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	db.AddVideoInfo(util.CreateUuid(), userId, title, url)

	c.JSON(http.StatusOK, model.Response{
		StatusCode: 0,
		StatusMsg:  "uploaded successfully",
	})
	return
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	token := c.Query("token")
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	logrus.Infof("[PublishList] token=%v, userId=%v", token, userId)
	if !service.CheckToken(userId, token) {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "check token failed",
		})
		return
	}

	videoList := service.GetPublishVideoList(userId)
	logrus.Infof("[PublishList] videoList=%+v", videoList)
	c.JSON(http.StatusOK, VideoListResponse{
		Response: model.Response{
			StatusCode: 0,
		},
		VideoList: videoList,
	})
	return
}

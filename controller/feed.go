package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/qingxunying/douyin/model"
	"github.com/qingxunying/douyin/service"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)

type FeedResponse struct {
	model.Response
	VideoList []model.Video `json:"video_list,omitempty"`
	NextTime  int64         `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	token := c.Query("token")
	currentTime := time.Now().Unix()
	latestTime, _ := strconv.ParseInt(c.DefaultQuery("latest_time", strconv.FormatInt(currentTime, 10)), 10, 64)
	var userId int64
	if token != "" {
		userId, _ = service.ParseToken(token)
	}
	videoList, nextTime := service.GetAllVideoList(userId, latestTime, currentTime)
	logrus.Infof("[Feed] videoList=%+v", videoList)
	c.JSON(http.StatusOK, FeedResponse{
		Response:  model.Response{StatusCode: 0},
		VideoList: videoList,
		NextTime:  nextTime,
	})
	return
}

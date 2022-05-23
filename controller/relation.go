package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/qingxunying/douyin/model"
	"github.com/qingxunying/douyin/service"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type UserListResponse struct {
	model.Response
	UserList []model.User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	token := c.Query("token")
	userId, userName := service.ParseToken(token)
	if userId == 0 || userName == "" {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	followUserId, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	actionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 64)
	service.AddFollowInfo(userId, followUserId, int(actionType))

	c.JSON(http.StatusOK, model.Response{StatusCode: 0})
	return
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	token := c.Query("token")
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if !service.CheckToken(userId, token) {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "check token failed",
		})
		return
	}

	followUserList := service.GetFollowUser(userId)
	logrus.Infof("[FollowList] followUserList=%+v", followUserList)
	c.JSON(http.StatusOK, UserListResponse{
		Response: model.Response{
			StatusCode: 0,
		},
		UserList: followUserList,
	})
	return
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	token := c.Query("token")
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if !service.CheckToken(userId, token) {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "check token failed",
		})
		return
	}

	followerUserList := service.GetFollowerUser(userId)
	logrus.Infof("[FollowerList] followerUserList=%+v", followerUserList)
	c.JSON(http.StatusOK, UserListResponse{
		Response: model.Response{
			StatusCode: 0,
		},
		UserList: followerUserList,
	})
	return
}

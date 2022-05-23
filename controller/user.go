package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/qingxunying/douyin/db"
	"github.com/qingxunying/douyin/model"
	"github.com/qingxunying/douyin/rdb"
	"github.com/qingxunying/douyin/service"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type UserLoginResponse struct {
	model.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	model.Response
	User model.User `json:"user"`
}

func Register(c *gin.Context) {
	userName := c.Query("username")
	password := c.Query("password")
	if userName == "" || password == "" {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "user_name or password empty"},
		})
		return
	}

	userInfo := db.GetUserInfoByUserName(userName)
	if userInfo != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
		return
	}
	userInfo, token := service.CreateUser(userName, password)
	if userInfo == nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "create userInfo error"},
		})
		return
	}
	logrus.Infof("[Register] userId=%+v, token=%+v", userInfo.UserId, token)
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: model.Response{StatusCode: 0},
		UserId:   userInfo.UserId,
		Token:    token,
	})
	return
}

func Login(c *gin.Context) {
	userName := c.Query("username")
	password := c.Query("password")
	if userName == "" || password == "" {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "user_name or password empty"},
		})
		return
	}

	userInfo := db.GetUserInfoByPassword(userName, password)
	if userInfo == nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
		return
	}
	token := rdb.GetToken(userInfo.UserId)
	logrus.Infof("[Login] userId=%+v, token=%+v", userInfo.UserId, token)
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: model.Response{StatusCode: 0},
		UserId:   userInfo.UserId,
		Token:    token,
	})
	return
}

func UserInfo(c *gin.Context) {
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	token := c.Query("token")
	userIdFromToken, _ := service.ParseToken(token)
	if userIdFromToken != userId {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "check token failed",
		})
		return
	}

	user := service.GetUser(userId, userId)
	if user == nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "User doesn't exist",
		})
		return
	}
	logrus.Infof("[UserInfo] user=%+v", *user)
	c.JSON(http.StatusOK, UserResponse{
		Response: model.Response{StatusCode: 0},
		User:     *user,
	})
	return
}

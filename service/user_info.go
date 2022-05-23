package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/qingxunying/douyin/constdef"
	"github.com/qingxunying/douyin/db"
	"github.com/qingxunying/douyin/model"
	"github.com/qingxunying/douyin/rdb"
	"github.com/qingxunying/douyin/util"
	"github.com/sirupsen/logrus"
	"sync"
)

func CreateUser(userName string, password string) (*db.UserInfo, string) {
	userId := util.CreateUuid()
	token := CreateToken(userId, userName)
	rdb.AddToken(userId, token)
	logrus.Infof("gen token=%v", token)
	return db.AddUserInfo(userId, userName, password), token
}

func CreateToken(userId int64, userName string) string {
	mapClaims := jwt.MapClaims{
		constdef.UserId:   userId,
		constdef.UserName: userName,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)
	tokenJson, _ := token.SignedString(rdb.GetRandomSalt())
	return tokenJson
}

func ParseToken(tokenJson string) (int64, string) {
	if tokenJson == "" {
		return 0, ""
	}
	salts := rdb.GetAllSalts()
	result := make(chan jwt.MapClaims, 1)
	var wg sync.WaitGroup
	for index := 0; index < len(salts); index++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			token, err := jwt.Parse(tokenJson, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(salts[index]), nil
			})
			if err == nil && token.Valid {
				claims, _ := token.Claims.(jwt.MapClaims)
				result <- claims

			}
		}(index)
	}
	wg.Wait()
	for item := range result {
		if item != nil {
			return int64(item[constdef.UserId].(float64)), item[constdef.UserName].(string)
		}
	}

	return 0, ""
}

func GetUser(userId int64, anotherUserId int64) *model.User {
	userInfo := db.GetUserInfoByUserId(userId)
	if userInfo == nil {
		logrus.Errorf("[GetUser] userInfo nil")
		return nil
	}
	isFollow := false
	if anotherUserId != 0 && (userId == anotherUserId || db.IsFollowedRelation(userId, anotherUserId)) {
		isFollow = true
	}
	followCount := db.GetFollowCount(userId)
	followerCount := db.GetFollowerCount(userId)
	user := &model.User{
		Id:            userInfo.UserId,
		Name:          userInfo.UserName,
		FollowCount:   followCount,
		FollowerCount: followerCount,
		IsFollow:      isFollow,
	}
	return user
}

func GetFollowUser(userId int64) []model.User {
	var followUsers []model.User
	followUserInfos := db.GetAllFollowUser(userId)
	for _, followUserInfo := range followUserInfos {
		followCount := db.GetFollowCount(followUserInfo.FollowUserId)
		followerCount := db.GetFollowerCount(followUserInfo.FollowUserId)
		name := db.GetUserNameByUserId(followUserInfo.FollowUserId)
		user := model.User{
			Id:            followUserInfo.UserId,
			Name:          name,
			FollowCount:   followCount,
			FollowerCount: followerCount,
			IsFollow:      true,
		}
		followUsers = append(followUsers, user)
	}
	return followUsers
}

func GetFollowerUser(userId int64) []model.User {
	var followerUsers []model.User
	followerUserInfos := db.GetAllFollowerUser(userId)
	for _, followerUserInfo := range followerUserInfos {
		followCount := db.GetFollowCount(followerUserInfo.FollowUserId)
		followerCount := db.GetFollowerCount(followerUserInfo.FollowUserId)
		name := db.GetUserNameByUserId(followerUserInfo.FollowUserId)
		isFollow := db.IsFollowedRelation(followerUserInfo.FollowUserId, userId)
		user := model.User{
			Id:            followerUserInfo.UserId,
			Name:          name,
			FollowCount:   followCount,
			FollowerCount: followerCount,
			IsFollow:      isFollow,
		}
		followerUsers = append(followerUsers, user)
	}
	return followerUsers
}

func CheckToken(userId int64, token string) bool {
	userIdFromToken, _ := ParseToken(token)
	if userId != userIdFromToken {
		return false
	}
	return true
}

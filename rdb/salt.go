package rdb

import (
	"github.com/qingxunying/douyin/constdef"
)

func GetAllSalts() []string {
	return rdb.SMembers(constdef.KeySalt).Val()
}

func GetRandomSalt() []byte {
	return []byte(rdb.SRandMemberN(constdef.KeySalt, 1).Val()[0])
}

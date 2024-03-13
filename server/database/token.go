package database

import (
	"blog/utils"
	"fmt"
	"time"
)

func SetToken(authToken string, refreshToken string) error {
	db := GetRedisConnection()
	if err := db.Set(refreshToken, authToken, 7*time.Hour*24).Err(); err != nil {
		utils.LogRus.Warnf("Redis 创建token失败")
		return err
	}
	return nil
}

func GetToken(refreshToken string) (string, error) {
	db := GetRedisConnection()
	res,err := db.Get(refreshToken).Result()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return "",err
	}
	return res,err
}

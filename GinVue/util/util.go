package util

import (
	"math/rand"
	"time"
)

// GetRandomName 返回一个十位的随机字符串
func GetRandomName() string {

	var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	result := make([]byte, 10)

	rand.Seed(time.Now().Unix())
	//随机化
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

package common

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateRandomString(length int) (string, error) {
	// 生成足够高熵的随机字节序列
	randomBytes := make([]byte, length)
	if _, err := rand.Read(randomBytes); err != nil {
		return "", err
	}

	// 使用base64.URLEncoding编码随机字节序列
	randomString := base64.URLEncoding.EncodeToString(randomBytes)

	return randomString[:length], nil
}

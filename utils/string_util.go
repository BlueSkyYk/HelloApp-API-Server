package utils

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"regexp"
	"time"
)

/**
	获取随机指定了长度的字符串
 */
func GetRandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

/**
	检查手机号
 */
func CheckPhone(phone string) bool {
	const regular = `^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`
	reg := regexp.MustCompile(regular)
	return reg.MatchString(phone)
}

/**
	获取字符MD5
 */
func GetStrMd5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

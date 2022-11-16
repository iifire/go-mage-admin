package helper

import (
	"math/rand"
	"time"
	"unsafe"
)

// String 字符串常用功能
type String struct {
}

// GetRandomString 生成指定长度的随机字符串
func (s *String) GetRandomString(n int) string {
	// GetRandomString 返回随机字符串
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var src = rand.NewSource(time.Now().UnixNano())
	const (
		// 6 bits to represent a letter index
		letterIdBits = 6
		// All 1-bits as many as letterIdBits
		letterIdMask = 1<<letterIdBits - 1
		letterIdMax  = 63 / letterIdBits
	)
	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdMax letters!
	for i, cache, remain := n-1, src.Int63(), letterIdMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdMax
		}
		if idx := int(cache & letterIdMask); idx < len(letters) {
			b[i] = letters[idx]
			i--
		}
		cache >>= letterIdBits
		remain--
	}
	return *(*string)(unsafe.Pointer(&b))
}

// Capitalize 字符串首字母转大写
func (s *String) Capitalize(str string) string {
	var upperStr string
	vv := []rune(str)
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 {
				vv[i] -= 32
				upperStr += string(vv[i])
			} else {
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}

// DeCapitalize 字符串首字母转小写
func (s *String) DeCapitalize(str string) string {
	var upperStr string
	vv := []rune(str)
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 65 && vv[i] <= 97 {
				vv[i] += 32
				upperStr += string(vv[i])
			} else {
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}

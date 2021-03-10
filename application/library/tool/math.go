package tool

import (
	"math/rand"
	"time"
)

// https://gist.github.com/DavadDi/8944292b59d7e74812cb91788218a246
var (
	codes   = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	codeLen = len(codes)
)

// 生成随机 Nonce
func RandNewStr(len int) string {
	data := make([]byte, len)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len; i++ {
		idx := rand.Intn(codeLen)
		data[i] = codes[idx]
	}

	return string(data)
}

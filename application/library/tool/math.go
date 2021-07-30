package tool

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
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

func KeepStringNum(value string, l int32) string {
	s := fmt.Sprintf("%%.%df", l)
	if v, err := strconv.ParseFloat(value, 64); err != nil {
		return "-"
	} else {
		return fmt.Sprintf(s, v)
	}
	//if d, e := decimal.NewFromString(value); e != nil {
	//	return "-"
	//} else {
	//	return d.Round(l).String()
	//}
}

func KeepFloatNum(value float64, l int32) string {
	s := fmt.Sprintf("%%.%df", l)
	return fmt.Sprintf(s, value)

	//return decimal.NewFromFloat(value).Round(l).String()
}


func FormatFloat(num float64, decimal int) string {
	// 默认乘1
	d := float64(1)
	if decimal > 0 {
		// 10的N次方
		d = math.Pow10(decimal)
	}
	// math.trunc作用就是返回浮点数的整数部分
	// 再除回去，小数点后无效的0也就不存在了
	return strconv.FormatFloat(math.Trunc(num*d)/d, 'f', -1, 64)
}
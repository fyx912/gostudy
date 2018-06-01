package common

import(
	"time"
	"strconv"
	"fmt"
	"crypto/md5"
)

// 生成uuid,根据当前系统时间的MD5值
func UUID() string{
	currentTime := time.Now().Unix()
	timedata := strconv.FormatInt(currentTime, 16)
	//%X    表示为十六进制，使用A-F
	md5Value := fmt.Sprintf("%X", md5.Sum([]byte(timedata)))
	return md5Value
}
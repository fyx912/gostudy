package common

import(
	"crypto/md5"
	// "encoding/hex"
	"fmt"
)

func MD5(data string)string{
	md5Byte := md5.Sum([]byte(data))
	md5Data := fmt.Sprintf("%X", md5Byte)
	return md5Data
}
/*工具包*/
package utils

import (
	"crypto/md5"
	"fmt"
)

/*获取md5加密字符串*/
func GetMd5String(str string) string{
	data := []byte(str)
	cipherStr := md5.Sum(data)
	md5str := fmt.Sprintf("%x", cipherStr)//将[]byte转成16进制
	return md5str
}
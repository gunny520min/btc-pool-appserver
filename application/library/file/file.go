package file

import (
	"os"
	"path"
	"strings"
)

// 创建目录
func CreateFileDir(filePath string) error {
	if !IsExist(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		return err
	}
	return nil
}

// 判断所给路径文件夹是否存在
func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

//获取文件扩展名
func GetFileExt(filepath string, inclueSep bool) string {
	ret := ""
	ret = path.Base(filepath)
	ret = path.Ext(filepath)
	if !inclueSep {
		ret = strings.TrimLeft(ret, ".")
	}
	return ret
}

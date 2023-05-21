package file

import (
	"fmt"
	"io/ioutil"
	"os"
)

//获取所有文件名,包含路径（相对路径）
func GetAllFile(pathname string) (s []string, err error) {
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		fmt.Println("read dir fail:", err)
		return s, err
	}
	for _, fi := range rd {
		if fi.IsDir() {
			fullDir := pathname + "/" + fi.Name()
			s, err = GetAllFile(fullDir)
			if err != nil {
				fmt.Println("read dir fail:", err)
				return s, err
			}
		} else {
			fullName := pathname + "/" + fi.Name()
			s = append(s, fullName)
		}
	}
	return s, nil
}

/*1、如果返回的错误为nil,说明文件或文件夹存在
2、如果返回的错误类型使用os.IsNotExist()判断为true,说明文件或文件夹不存在
3、如果返回的错误为其它类型,则不确定是否在存在*/
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// 判断所给路径是否为文件夹

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 判断所给路径是否为文件

func IsFile(path string) bool {
	return !IsDir(path)
}

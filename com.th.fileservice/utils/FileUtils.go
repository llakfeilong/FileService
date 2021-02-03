package utils

import (
	"io/ioutil"
	"log"
	"os"
)

func GetFiles(folder string) []string {
	filespath := make([]string, 0)
	files, err := ioutil.ReadDir(folder) //specify the current dir
	if err != nil {
		log.Println("打开文件夹发生错误:", err)
	}
	for _, file := range files {
		if file.IsDir() {
			GetFiles(folder + "/" + file.Name())
		} else {
			log.Println(folder + file.Name())
			filespath = append(filespath, folder+file.Name())
		}
	}
	return filespath
}

func CreateFilefolder(path string) {
	err := os.MkdirAll(path, 0766)
	if err != nil {
		log.Println(err)
	}
}

/*
   判断文件或文件夹是否存在
   如果返回的错误为nil,说明文件或文件夹存在
   如果返回的错误类型使用os.IsNotExist()判断为true,说明文件或文件夹不存在
   如果返回的错误为其它类型,则不确定是否在存在
*/
func FilePathExists(path string) (bool, error) {

	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

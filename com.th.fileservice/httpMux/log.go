package httpMux

import (
	"log"
	"os"
	"time"
)

func formatTime() string {
	return time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")
}

/**
 * 初始化日志配置
 */
func InitLog(path string, name string, level string) {
	err := os.MkdirAll(path, 0766)
	if err != nil {
		log.Println(err)
	}
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	if level == "prod" {
		// 获取日志文件句柄
		// 已 只写入文件|没有时创建|文件尾部追加 的形式打开这个文件
		logFile, err := os.OpenFile(path+name, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
		// 设置存储位置
		log.SetOutput(logFile)
	}
}

//func Info(value ...interface{})  {
//	log.Println("[INFO]",value)
//}
//
//func Err(value ...interface{})  {
//	log.Println("[ERROR]",value)
//}

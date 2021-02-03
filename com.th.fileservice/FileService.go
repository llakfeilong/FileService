package main

import (
	"com.th.fileservice/config"
	"com.th.fileservice/controller"
	"com.th.fileservice/dbManger"
	"com.th.fileservice/httpMux"
	"com.th.fileservice/nacos"
	"log"
)

func main() {
	//读取配置
	config.LoadConfig()
	//初始化数据库
	dbManger.InitDB()
	//nacos
	nacos.Run()
	//httpmux
	router := httpMux.NewMux()
	router.GET("/download/getFileByteArray", controller.GetFileByteArray)
	router.GET("/v1/download/getFileBase64", controller.GetFileBase64)
	router.POST("/uploadFile", controller.UploadFile)
	router.POST("/v1/uploadFile", controller.UploadFileToFileType)
	router.POST("/v1/batchDeleteFiles", controller.BatchDeleteFile)
	router.POST("/fileManger/addFileType", controller.AddFileType)
	router.POST("/fileManger/updateFileType", controller.UpdateFileType)
	router.POST("/fileManger/pageList", controller.PageList)
	log.Println("###http启动端口:" + config.GetWebConfig().Port)
	log.Println("###数据库URL:" + config.GetDBConfig().Address)
	log.Println("###nacos reg:"+config.GetNacosCfg().RegisterCfg.Ip, config.GetNacosCfg().RegisterCfg.ServiceName, "service ip:", config.GetNacosCfg().ServerCfg.IpAddr)
	router.Run(":" + config.GetWebConfig().Port)

}

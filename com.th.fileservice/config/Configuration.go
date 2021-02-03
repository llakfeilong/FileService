package config

import (
	"com.th.fileservice/httpMux"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"runtime"
	"strconv"
)

var filePath = "config.yml"

var gormCfg = GormCfg{}

var storageCfg = StorageCfg{}

var nacosCfg = NacosCfg{}

var webCfg = WebCfg{}

var logCfg = LogCfg{}

type GormCfg struct {
	Username         string `mapstructure:"username"`
	Password         string `mapstructure:"password"`
	Dbname           string `mapstructure:"dbname"`
	Address          string `mapstructure:"address"`
	Sqlinitpath      string `mapstructure:"sqlinitpath"`
	Sqlincrementpath string `mapstructure:"sqlincrementpath"`
}

type StorageCfg struct {
	WinStorage   string
	LinuxStorage string
}

type ClientConfig struct {
	NamespaceId string
	TimeoutMs   uint64
	LogDir      string
	CacheDir    string
	LogLevel    string
}

type ServerConfig struct {
	IpAddr string
	Port   uint64
}

type RegisterConfig struct {
	Ip          string
	ServiceName string
	ClusterName string
	GroupName   string
	Weight      string
}

type NacosCfg struct {
	ClientCfg   ClientConfig
	ServerCfg   ServerConfig
	RegisterCfg RegisterConfig
}

type LogCfg struct {
	Path  string
	Level string
	Name  string
}

type WebCfg struct {
	Port string
}

func LoadConfig() {
	// 设置配置文件信息
	viper.SetConfigType("yml")
	viper.SetConfigFile(filePath)

	// 读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("读取配置文件失败, 异常信息 : ", err)
	}

	//// 直接从viper对象中获取key的value数据,并且可以定义类型
	//fmt.Println(viper.Get("user.name"))
	//fmt.Println(viper.GetInt("user.age"))
	//
	//// 判断key是否存在, 返回true/false
	//fmt.Println(viper.IsSet("user.name"))
	//
	//// 设置key的value值, 优先级最高, 可以在读取配置文件之前设置
	//viper.Set("user.age",19)
	//日志配置
	logCfg.Path = fmt.Sprintf("%v", viper.Get("log.path"))
	logCfg.Level = fmt.Sprintf("%v", viper.Get("log.level"))
	logCfg.Name = fmt.Sprintf("%v", viper.Get("log.name"))
	httpMux.InitLog(logCfg.Path, logCfg.Name, logCfg.Level)
	//web配置
	webCfg.Port = fmt.Sprintf("%v", viper.Get("web.port"))
	//数据库配置
	gormCfg.Sqlinitpath = fmt.Sprintf("%v", viper.Get("gorm.sqlinitpath"))
	gormCfg.Sqlincrementpath = fmt.Sprintf("%v", viper.Get("gorm.sqlincrementpath"))
	gormCfg.Dbname = fmt.Sprintf("%v", viper.Get("gorm.dbname"))
	gormCfg.Address = fmt.Sprintf("%v", viper.Get("gorm.address"))
	gormCfg.Username = fmt.Sprintf("%v", viper.Get("gorm.username"))
	gormCfg.Password = fmt.Sprintf("%v", viper.Get("gorm.password"))

	//文件存储基本路径配置
	storageCfg.LinuxStorage = fmt.Sprintf("%v", viper.Get("storagePath.linuxStorage"))
	storageCfg.WinStorage = fmt.Sprintf("%v", viper.Get("storagePath.winStorage"))

	//nacos配置
	//客户端配置
	nacosCfg.ClientCfg.NamespaceId = fmt.Sprintf("%v", viper.Get("nacos.clientConfig.NamespaceId"))
	timeout, _ := strconv.Atoi(fmt.Sprintf("%v", viper.Get("nacos.clientConfig.TimeoutMs")))
	nacosCfg.ClientCfg.TimeoutMs = uint64(timeout)
	nacosCfg.ClientCfg.CacheDir = fmt.Sprintf("%v", viper.Get("nacos.clientConfig.CacheDir"))
	nacosCfg.ClientCfg.LogDir = fmt.Sprintf("%v", viper.Get("nacos.clientConfig.LogDir"))
	nacosCfg.ClientCfg.LogLevel = fmt.Sprintf("%v", viper.Get("nacos.clientConfig.LogLevel"))
	//服务端配置
	nacosCfg.ServerCfg.IpAddr = fmt.Sprintf("%v", viper.Get("nacos.serverConfig.IpAddr"))
	nacosport, _ := strconv.Atoi(fmt.Sprintf("%v", viper.Get("nacos.serverConfig.Port")))
	nacosCfg.ServerCfg.Port = uint64(nacosport)
	//注册配置
	nacosCfg.RegisterCfg.Ip = fmt.Sprintf("%v", viper.Get("nacos.registerConfig.Ip"))
	nacosCfg.RegisterCfg.ServiceName = fmt.Sprintf("%v", viper.Get("nacos.registerConfig.ServiceName"))
	nacosCfg.RegisterCfg.ClusterName = fmt.Sprintf("%v", viper.Get("nacos.registerConfig.ClusterName"))
	nacosCfg.RegisterCfg.GroupName = fmt.Sprintf("%v", viper.Get("nacos.registerConfig.GroupName"))
	nacosCfg.RegisterCfg.Weight = fmt.Sprintf("%v", viper.Get("nacos.registerConfig.Weight"))

	// 将文件内容解析后封装到cfg对象中
	//err = viper.Unmarshal(&gormCfg)
	//if err != nil {
	//	fmt.Println("解析配置文件失败, 异常信息 : ", err)
	//}
	log.Println("配置加载完成")
}

// 使用时直接调用该方法即可
func GetDBConfig() GormCfg {
	return gormCfg
}

func GetNacosCfg() NacosCfg {
	return nacosCfg
}

func GetWebConfig() WebCfg {
	return webCfg
}

func GetLogConfig() LogCfg {
	return logCfg
}

//判断当前操作系统
func GetStoragePath() string {
	switch runtime.GOOS {
	case "windows":
		return storageCfg.WinStorage
	case "linux":
		return storageCfg.LinuxStorage
	}
	return ""
}

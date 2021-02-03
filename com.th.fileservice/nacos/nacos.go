package nacos

import (
	"com.th.fileservice/config"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"log"
	"strconv"
)

func Run() {
	//创建clientConfig
	clientConfig := constant.ClientConfig{
		NamespaceId:         config.GetNacosCfg().ClientCfg.NamespaceId, // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId
		TimeoutMs:           config.GetNacosCfg().ClientCfg.TimeoutMs,
		NotLoadCacheAtStart: true,
		LogDir:              config.GetNacosCfg().ClientCfg.LogDir,
		CacheDir:            config.GetNacosCfg().ClientCfg.CacheDir,
		RotateTime:          "1h",
		MaxAge:              3,
		LogLevel:            config.GetNacosCfg().ClientCfg.LogLevel,
	}

	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      config.GetNacosCfg().ServerCfg.IpAddr,
			ContextPath: "/nacos",
			Port:        config.GetNacosCfg().ServerCfg.Port,
			Scheme:      "http",
		},
	}

	// 创建服务发现客户端
	namingClient, err := clients.CreateNamingClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		fmt.Println(err)
	}

	// 创建动态配置客户端
	//configClient, err := clients.CreateConfigClient(map[string]interface{}{
	//	"serverConfigs": serverConfigs,
	//	"clientConfig":  clientConfig,
	//})
	//if err != nil{
	//	fmt.Println(err)
	//}

	port, _ := strconv.ParseInt(config.GetWebConfig().Port, 10, 64)
	weight, _ := strconv.ParseFloat(config.GetNacosCfg().RegisterCfg.Weight, 64)
	namingsuccess, err := namingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          config.GetNacosCfg().RegisterCfg.Ip,
		Port:        uint64(port),
		ServiceName: config.GetNacosCfg().RegisterCfg.ServiceName,
		Weight:      weight,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    map[string]string{"preserved.register.source": "SPRING_CLOUD"},
		ClusterName: config.GetNacosCfg().RegisterCfg.ClusterName, // 默认值DEFAULT
		GroupName:   config.GetNacosCfg().RegisterCfg.GroupName,   // 默认值DEFAULT_GROUP
	})

	//configsuccess, err := configClient.PublishConfig(vo.ConfigParam{
	//	DataId:  "dataId",
	//	Group:   "group",
	//	Content: "hello world!222222"})
	if namingsuccess {
		log.Println("nacos注册成功")
	} else {
		log.Println("nacos注册失败")
	}
}

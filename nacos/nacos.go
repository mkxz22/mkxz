package nacos

import (
	"encoding/json"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

type T struct {
	Mysql struct {
		Addr     string `json:"addr"`
		Port     int    `json:"port"`
		Password string `json:"password"`
		Data     string `json:"data"`
		User     string `json:"user"`
	} `json:"mysql"`
	Redis struct {
		Addr     string `json:"addr"`
		Password string `json:"password"`
		Db       int    `json:"db"`
	} `json:"redis"`
	Elasticsearch struct {
		Addr string `json:"addr"`
	} `json:"elasticsearch"`
	Aliyun struct {
		AccessKeyid     string `json:"accessKeyid"`
		Accesskeysecret string `json:"accesskeysecret"`
		Region          string `json:"region"`
		Bucketname      string `json:"bucketname"`
	} `json:"aliyun"`
}

var ns T

func Naocs(namespaceId, IpAddr, DataId, Group string, Port int) {
	//create clientConfig
	clientConfig := constant.ClientConfig{
		NamespaceId:         namespaceId, //we can create multiple clients with different namespaceId to support multiple namespace.When namespace is public, fill in the blank string here.
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}
	// At least one ServerConfig
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      IpAddr,
			ContextPath: "/nacos",
			Port:        uint64(Port),
			Scheme:      "http",
		},
	}
	// Create config client for dynamic configuration
	clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	// Another way of create config client for dynamic configuration (recommend)
	configClient, _ := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	content, _ := configClient.GetConfig(vo.ConfigParam{
		DataId: DataId,
		Group:  Group})
	err := json.Unmarshal([]byte(content), &ns)
	if err != nil {
		return
	}
	fmt.Println(content)
}

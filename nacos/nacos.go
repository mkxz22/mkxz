package nacos

import (
	"encoding/json"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

type Nacos struct {
	NamespaceId string
	Group       string
	DataId      string
	Port        int
	IpAddr      string
}
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

var (
	Configs  T
	NacosSrv Nacos
)

func NacosInit(namespaceId, ipAddr, dataId, group string, port uint64) (T, Nacos) {
	clientConfig := constant.ClientConfig{
		NamespaceId:         namespaceId, // 如果需要支持多namespace，我们可以创建多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      ipAddr,
			ContextPath: "/nacos",
			Port:        port,
			Scheme:      "http",
		},
	}
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		return T{}, NacosSrv
	}
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group})
	configAdd := T{}
	json.Unmarshal([]byte(content), &configAdd)
	Configs = configAdd
	return Configs, NacosSrv
}

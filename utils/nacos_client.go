package utils

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/cobra"
	"io"
	"text/template"
)

type NacosClient struct {
	client config_client.IConfigClient
}

func NewNacosClient(addr, namespace string, port uint64) *NacosClient {
	serviceConfig := []constant.ServerConfig{
		{
			IpAddr: addr,
			Port:   port,
		},
	}

	// 创建clientConfig
	clientConfig := constant.ClientConfig{
		TimeoutMs:           5000,
		NamespaceId:         namespace,
		CacheDir:            "/tmp/nacos/cache",
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/logs",
		LogLevel:            "debug",
		AppendToStdout:      false,
	}

	// 创建动态配置客户端
	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serviceConfig,
		},
	)

	cobra.CheckErr(err)

	return &NacosClient{client: configClient}
}


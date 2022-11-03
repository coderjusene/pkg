package nacos

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/app/middlewares/client/sd"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/registry/nacos"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

type Client struct {
	client *client.Client
}

func NewClient(addr string, port uint64, namespace string) *Client {
	cli, err := client.NewClient()
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(addr, port),
	}

	cc := constant.ClientConfig{
		NamespaceId:         namespace,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		LogLevel:            "info",
	}

	nacosCli, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)

	if err != nil {
		hlog.Fatal(err)
	}

	r := nacos.NewNacosResolver(nacosCli)
	cli.Use(sd.Discovery(r))

	return &Client{cli}
}

func (c *Client) Get(ds, path string, headers map[string]string) (status int, body []byte, err error) {
	url := fmt.Sprintf("http://%s%s", ds, path)
	req := &protocol.Request{}
	res := &protocol.Response{}
	req.SetMethod(consts.MethodGet)
	req.SetRequestURI(url)
	req.SetHeaders(headers)
	option := req.Options()
	requestOptions := []config.RequestOption{}
	requestOptions = append(requestOptions, config.WithSD(true))
	option.Apply(requestOptions)
	err = c.client.Do(context.Background(), req, res)
	return res.StatusCode(), res.Body(), err
}

func (c *Client) GetWithQuery(ds, path string, headers map[string]string, query string) (status int, body []byte, err error) {
	url := fmt.Sprintf("http://%s%s", ds, path)
	req := &protocol.Request{}
	res := &protocol.Response{}
	req.SetMethod(consts.MethodGet)
	req.SetRequestURI(url)
	req.SetHeaders(headers)
	req.SetQueryString(query)
	option := req.Options()
	requestOptions := []config.RequestOption{}
	requestOptions = append(requestOptions, config.WithSD(true))
	option.Apply(requestOptions)
	err = c.client.Do(context.Background(), req, res)
	return res.StatusCode(), res.Body(), err
}

func (c *Client) Post(ds, path string) {

}
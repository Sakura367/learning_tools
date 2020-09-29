package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/client/grpc"
	"github.com/micro/go-micro/v2/client/selector"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	test_agent "micro_v2"
	"os"
	"time"
)

func main() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	micReg := etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))
	agent := test_agent.NewTestService("srv.test.client", grpc.NewClient(
		client.Registry(micReg),
		client.Selector(selector.NewSelector(
			selector.Registry(micReg),
		)),
	))

	var opss client.CallOption = func(o *client.CallOptions) {
		o.RequestTimeout = time.Second * 30
		o.DialTimeout = time.Second * 30
		o.Retries = 3
		o.Address = []string{"127.0.0.1:8081"}
	}
	info, err := agent.RpcUserInfo(context.TODO(), &test_agent.ReqMsg{
		UserName: "test user",
	}, opss)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Println(info)
}
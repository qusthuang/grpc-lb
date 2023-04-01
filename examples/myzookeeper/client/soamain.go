package main

import (
	"github.com/liyue201/grpc-lb/balancer"
	"github.com/liyue201/grpc-lb/examples/proto3"
	registry "github.com/liyue201/grpc-lb/registry/zookeeper"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	registry.RegisterResolver2("zk", []string{"localhost:2181"}, "/soav2/HelloSoaTest/fat/providerList")
	c, err := grpc.Dial("zk:///", grpc.WithInsecure(), grpc.WithBalancerName(balancer.RoundRobin))
	if err != nil {
		log.Printf("grpc dial: %s", err)
		return
	}
	defer c.Close()
	client := proto3.NewSoaInvokerServiceClient(c)

	for i := 0; i < 5; i++ {
		// string reqId = 1;
		//    string rpcId = 2;
		//    string iface = 3;
		//    string method = 4;
		//    string requestJson = 5;
		resp, err := client.Call(context.Background(), &proto3.SoaInvokerRequest{ReqId: "1", RpcId: "11", Iface: "org.example.TestIface", Method: "execute", RequestJson: ""})
		if err != nil {
			log.Println("aa:", err)
			time.Sleep(time.Second)
			continue
		}
		//time.Sleep(time.Second)
		log.Println(resp.Code)
	}
}

package main

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	studentdemo "github.com/han-y-s/kitex_demo/kitex_gen/student_demo/studentservice"
	consul2 "github.com/kitex-contrib/config-consul/consul"
	consulserver "github.com/kitex-contrib/config-consul/server"
	consul "github.com/kitex-contrib/registry-consul"
	"log"
	"net"
)

var (
	severList = []string{":9998"}
)

func main() {
	var err error
	r, err := consul.NewConsulRegister("127.0.0.1:8500") // r should not be reused.
	if err != nil {
		log.Fatal(err)
	}
	serviceName := "student_demo"
	consulClient, _ := consul2.NewClient(consul2.Options{})
	for _, s := range severList {
		go func(s string) {
			addr, _ := net.ResolveTCPAddr("tcp", s)
			svr := studentdemo.NewServer(new(StudentServiceImpl),
				server.WithRegistry(r),
				server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: serviceName}),
				server.WithServiceAddr(addr),
				server.WithSuite(consulserver.NewSuite(serviceName, consulClient)),
			)

			err = svr.Run()
			if err != nil {
				log.Println(err.Error())
			}
		}(s)
	}
	var c chan struct{} //nil channel
	<-c
}

//etcd
//func main() {
//	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"}) // r should not be reused.
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	for _, s := range severList {
//		go func(s string) {
//			addr, _ := net.ResolveTCPAddr("tcp", s)
//			svr := student_demo.NewServer(new(StudentServiceImpl), server.WithRegistry(r),
//				server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "student_demo"}),
//				server.WithServiceAddr(addr))
//
//			err = svr.Run()
//			if err != nil {
//				log.Println(err.Error())
//			}
//		}(s)
//	}
//	var c chan struct{} //nil channel
//	<-c
//}

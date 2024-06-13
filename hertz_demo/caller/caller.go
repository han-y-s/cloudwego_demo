package caller

import (
	"github.com/cloudwego/kitex/client"
	bgeneric "github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/han-y-s/hertz_demo/constants"
	"github.com/han-y-s/hertz_demo/kitex_gen/student_demo/studentservice"
	consulclient "github.com/kitex-contrib/config-consul/client"
	consul2 "github.com/kitex-contrib/config-consul/consul"
	"github.com/kitex-contrib/config-consul/utils"
	consul "github.com/kitex-contrib/registry-consul"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var Rpccli studentservice.Client
var GeneralCli bgeneric.Client

type configLog struct{}

func (cl *configLog) Apply(opt *utils.Options) {
	fn := func(k *consul2.Key) {
		klog.Infof("consul config %v", k)
	}
	opt.ConsulCustomFunctions = append(opt.ConsulCustomFunctions, fn)
}

var Provider *generic.ThriftContentProvider

// consul热更新泛化
func GeneralCliInit() {

	var err error

	consulClient, err := consul2.NewClient(consul2.Options{})
	if err != nil {
		panic(err)
	}
	cl := &configLog{}
	serviceName := "student_demo"
	clientName := "student_demo_client"

	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		panic(err)
	}
	Provider, err = generic.NewThriftContentProviderWithDynamicGo(constants.OldIdlContent, map[string]string{})
	if err != nil {
		panic(err)
	}
	// 构造http类型的泛化调用
	g, err := generic.HTTPThriftGeneric(Provider)
	if err != nil {
		panic(err)
	}
	GeneralCli, err = bgeneric.NewClient(
		serviceName,
		g,
		client.WithResolver(r),
		client.WithSuite(consulclient.NewSuite(serviceName, clientName, consulClient, cl)),
	)
	if err != nil {
		panic(err)
	}

}

//// etcd热更新泛化
//func GeneralCliInit() {
//
//	var err error
//	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
//	if err != nil {
//		panic(err)
//	}
//	Provider, err = generic.NewThriftContentProviderWithDynamicGo(constants.OldIdlContent, map[string]string{})
//	if err != nil {
//		panic(err)
//	}
//	// 构造http类型的泛化调用
//	g, err := generic.HTTPThriftGeneric(Provider)
//	if err != nil {
//		panic(err)
//	}
//	GeneralCli, err = bgeneric.NewClient("student_demo", g, client.WithResolver(r))
//	if err != nil {
//		panic(err)
//	}
//
//}

////普通泛化
//func GeneralCliInit() {
//	var err error
//	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
//	if err != nil {
//		panic(err)
//	}
//	p, err := generic.NewThriftFileProvider("/Users/bytedance/go/src/code.byted.org/pku/hertz_demo/idl/student_demo.thrift")
//	if err != nil {
//		panic(err)
//	}
//	// 构造http类型的泛化调用
//	g, err := generic.HTTPThriftGeneric(p)
//	if err != nil {
//		panic(err)
//	}
//	GeneralCli, err = bgeneric.NewClient("student_demo", g, client.WithResolver(r))
//	if err != nil {
//		panic(err)
//	}
//
//}

func ClientInit() {
	var err error

	//consulClient, err := consul2.NewClient(consul2.Options{})
	//if err != nil {
	//	panic(err)
	//}
	//cl := &configLog{}
	serviceName := "student_demo"
	//clientName := "student_demo_client"

	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		panic(err)
	}

	Rpccli, err = studentservice.NewClient(
		serviceName,
		client.WithResolver(r),
		//client.WithSuite(consulclient.NewSuite(serviceName, clientName, consulClient, cl)),
	)
	if err != nil {
		panic(err)
	}
}

package main

import (
	"Tiktok/app/order/handler"
	"Tiktok/app/order/idl/gen"
	"Tiktok/app/order/initialize"
	"flag"
	"fmt"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	initialize.InitiallizeLogger()
	initialize.InitializeConfig()
	initialize.InitializeDB()
	initialize.InitETCD()
	//initialize.InitializeSrv()

	IP := flag.String("ip","0.0.0.0","ip地址")
	Port := flag.Int("port",0,"端口号")
	flag.Parse()
	zap.S().Info("ip: ",*IP)
	grpcServer := grpc.NewServer()
	orderSrv := handler.OrderService{}
	gen.RegisterOrderServiceServer(grpcServer,&orderSrv)
	ServiceAddr := fmt.Sprintf("%s:%d",*IP,*Port)
	listener, err := net.Listen("tcp",ServiceAddr)
	if err!= nil {
		zap.S().Panic("failed to listen:" + err.Error())	
	}
	zap.S().Info("Service is listening on ",ServiceAddr)
	
	err = initialize.RegisterService(ServiceAddr)
	if err != nil {
	   zap.S().Info("注册服务失败")	
	}

	if err := grpcServer.Serve(listener); err != nil {
		zap.S().Panic("failed to start grpc:" + err.Error())
	}
}
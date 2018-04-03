package main

import (
	"bmsg/models"
	"bmsg/controllers"
	"bmsg/logger"
	_ "bmsg/routers"
	_ "github.com/lib/pq"
	"sync"

	pb "bmsg/protos"
	"bmsg/config"
	"github.com/astaxie/beego"
	"fmt"
	"net"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func init() {
	beego.BConfig.AppName = config.AppConf.AppName
	beego.BConfig.RunMode = config.AppConf.RunMode
	beego.BConfig.CopyRequestBody = config.AppConf.CopyRequestBody
	beego.BConfig.Listen.HTTPPort = config.AppConf.HttpPort
	beego.BConfig.WebConfig.EnableDocs = config.AppConf.EnableDocs
	beego.BConfig.WebConfig.AutoRender = config.AppConf.AutoRender
	models.InitDB()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		logger.Debug("grpc server start")
		grpcStart()
		logger.Debug("grpc done")
		wg.Done()
	}()
	go func() {
		logger.Debug("http server start")
		beego.Run()
		logger.Debug("http server done")
		wg.Done()
	}()

	wg.Wait()
	logger.Warn("Ser done")
}

func grpcStart() {
	lis, err := net.Listen("tcp", config.AppConf.GrpcListen)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	fmt.Printf("Message servic grpc-port(%v)\n", config.AppConf.GrpcListen)
	pb.RegisterMessageServiceServer(s, &controllers.MsgSerController{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
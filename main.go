package main

import (
	"bmsg/models"
	"bmsg/controllers"
	_ "bmsg/routers"
	_ "github.com/lib/pq"

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
	//if beego.BConfig.RunMode == "dev" {
	//	beego.BConfig.WebConfig.DirectoryIndex = true
	//	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	//}
	//beego.Run()
	grpcStart()
}

func grpcStart() {
	lis, err := net.Listen("tcp", config.AppConf.GrpcListen)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	fmt.Printf("Jwt servic grpc-port(%v)\n", config.AppConf.GrpcListen)
	//grpc token
	pb.RegisterMessageServiceServer(s, &controllers.MsgSerController{})
	//grpc email
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
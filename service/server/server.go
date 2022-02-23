package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/jordan-wright/email"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"mail2/service/common"
	"net"
	"net/smtp"
	"time"
	//"test/grpc_test/common"
)

const (
	Address = "10.243.105.17:8181"
)

var (
	port = flag.Int("port", 8181, "the server port")
)

type server struct {
	common.UnimplementedEmailServiceServer
}

func (s *server) SendEmail(ctx context.Context, request *common.GetEmailRequest) (*common.GetEmailResponse, error) {
	timestamp := request.Timestamp //报警时间搓
	moment := time.Unix(timestamp, 0).Format("2006-01-02  15:04:05")
	sender := request.Sender
	recipient := request.Recipient
	cpuUsed := request.CpuUsed   //cpu使用情况
	memUsed := request.MemUsed   //内存使用情况
	diskUsed := request.DiskUsed //磁盘使用情况
	grade := request.Grade       //报警等级
	// 简单设置 log 参数
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	em := email.NewEmail()
	// 设置 sender 发送方 的邮箱
	em.From = sender
	// 设置 receiver 接收方 的邮箱
	em.To = []string{recipient}

	// 发送内容
	em.Subject = fmt.Sprintf("Warning:The system monitoring and forecast level is %v", grade)
	// 简单设置文件发送的内容，暂时设置成纯文本
	content := fmt.Sprintf("%v : The occupancy of the system cpu is:%v ;"+
		" the occupancy of the system memory is %v ，the occupancy of the disk is %v", moment, cpuUsed, memUsed, diskUsed)
	em.Text = []byte(content)

	//设置服务器相关的配置
	err := em.Send("smtp.qq.com:25", smtp.PlainAuth("", "1834960035@qq.com", "ruomvvskfwpxdccb", "smtp.qq.com"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(sender + "  sent to " + recipient + " successfully ... ")

	return &common.GetEmailResponse{Code: 0, Info: "sending success"}, nil
}

/*func (s *server) mustEmbedUnimplementedEmailServiceServer() {

	panic("implement me")
}*/

/*func ConcurrentSend() {

}*/

func main() {
	listen, err := net.Listen("tcp", fmt.Sprintf("10.243.105.17:%d", *port))
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	common.RegisterEmailServiceServer(s, &server{})
	// 注册反射服务 这个服务是CLI使用的 跟服务本身没有关系
	reflection.Register(s)
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to server: %v", err)
	}

}

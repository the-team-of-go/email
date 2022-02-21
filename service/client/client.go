package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"mail2/service/common"
	//"test/grpc_test/common"
	"time"

	"google.golang.org/grpc"
)

var (
	//addr = flag.String("addr", "localhost:8181", "the address to connect to")
	addr = flag.String("addr", "localhost:8181", "the address to connect to")
)

func main() {
	//建立链接
	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := common.NewEmailServiceClient(conn)

	// Contact the server and print out its response.
	/*var id int64 = 1
	if len(os.Args) > 1 {
		id, _ = strconv.ParseInt(os.Args[1], 10, 64)
	}*/
	// 1秒的上下文
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SendEmail(ctx, &common.GetEmailRequest{
		Sender:    "1834960035@qq.com",
		Recipient: "stackoverflow520@163.com",
		CpuUsed:   "%30",
		MemUsed:   "%40",
		Grade:     "serious",
	})
	if err != nil {
		log.Fatalf("could not sendEmail: %v", err)
	}
	fmt.Println(r.Code)
	fmt.Println(r.Info)
}

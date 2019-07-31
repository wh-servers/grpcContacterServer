package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/gogo/protobuf/proto"
	pb "github.com/wh-servers/grpcContacterServer/contacter"
	"google.golang.org/grpc"
)

var (
	configFile = flag.String("conf", "conf/contacterDB.xml", "conf file address")
	addr       = flag.String("addr", "localhost:8081", "tcp address for server")
)

func main() {
	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		fmt.Println("dial err: ", err)
		return
	}
	client := pb.NewContacterAPIClient(conn)
	req := &pb.DemoRequest{}
	req.DemoInput = proto.String("hi , i am the input string")
	res := &pb.DemoResponse{}
	res, err = client.DemoAPI(context.Background(), req)
	if err != nil {
		fmt.Println("request err: ", err)
		return
	}
	fmt.Println(res)
}

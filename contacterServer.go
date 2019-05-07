package main

import (
	"flag"
	"net"

	service "github.com/wh-servers/grpcContacterServer/Service"
	pb "github.com/wh-servers/grpcContacterServer/contacter"
	db "github.com/wh-servers/grpcContacterServer/db/service"
	"google.golang.org/grpc"
)

var (
	configFile = flag.String("conf", "conf/contacterDB.xml", "conf file address")
	addr       = flag.String("addr", "localhost:80", "tcp address for server")
)

func main() {
	flag.Parse()
	err := db.OrmInit(*configFile)
	if err != nil {
		panic(err)
	}
	go db.LoadMapRun()
	serveRpc()
	//todo: shut down gracefully.
}

func serveRpc() {
	listener, err := net.Listen("tcp", *addr)
	if err != nil {
		panic(err)
	}
	contacterServer := service.ContacterServer{}
	srv := grpc.NewServer()
	pb.RegisterContacterAPIServer(srv, contacterServer) //contacterServer implemented all functions in ContacterApiServer(interface), so can use here.
	err = srv.Serve(listener)
	if err != nil {
		panic(err)
	}
}

package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	// port = []string{"1234", "2234"}
	port := flag.Int("port", 1234, "server port")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	stringService := new(StringService)
	stringService.Port = *port

	pb.RegisterStringServiceServer(grpcServer, stringService)
	grpcServer.Serve(lis)
}

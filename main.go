package main

import (
	pb "github.com/2014BDuck/tag-service/proto"
	"github.com/2014BDuck/tag-service/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

var port = "8002"

func main() {
	s := grpc.NewServer()
	pb.RegisterTagServiceServer(s, server.NewTagServer())

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	log.Printf("Listening: %s", port)
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("server.Server err: %v", err)
	}
}

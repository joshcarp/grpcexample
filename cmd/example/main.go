package main

import (
	"github.com/joshcarp/grpcexample/pkg/example"
	"github.com/joshcarp/grpcexample/proto/examplepb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	srv := grpc.NewServer()
	examplepb.RegisterFooAPIServer(srv, example.FooServer{})
	examplepb.RegisterBarAPIServer(srv, example.BarServer{})
	reflection.Register(srv)
	err = srv.Serve(ln)
	if err != nil {
		log.Fatal(err)
	}
}
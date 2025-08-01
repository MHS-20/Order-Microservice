package grpc

import (
	"fmt"
	"net"

	"log"

	"github.com/MHS-20/Order-Microservice/order/config"
	"github.com/MHS-20/Order-Microservice/order/internal/ports"
	"github.com/MHS-20/Order-Proto/golang/order"
	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
)

type Adapter struct {
	api    ports.APIPort
	port   int
	server *grpc.Server
	order.UnimplementedOrderServer
}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{api: api, port: port}
}

func (a Adapter) Run() {
	var err error
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen on port %d, error: %v", a.port, err)
	}

	grpcServer := grpc.NewServer()
	order.RegisterOrderServer(grpcServer, a)
	if config.GetEnv() == "development" {
		reflection.Register(grpcServer)
	}
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve grpc on port %d, error: %v", a.port, err)
	}
}

func (a Adapter) Stop() {
	a.server.Stop()
}

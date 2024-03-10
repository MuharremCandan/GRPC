package gapi

import (
	"log"
	"net"
	"test-grpc-project/pkg/config"
	"test-grpc-project/pkg/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
)

type Server struct {
	pb.UnimplementedGrpcProjectServer
	db     *gorm.DB
	config *config.Config
}

func NewServer(config *config.Config, db *gorm.DB) *Server {
	return &Server{
		config: config,
		db:     db,
	}
}

func (s *Server) NewgRpcServer() {
	grpcServer := grpc.NewServer()
	pb.RegisterGrpcProjectServer(grpcServer, s)
	reflection.Register(grpcServer)
	listener, err := net.Listen("tcp", net.JoinHostPort(s.config.GrpcServer.Host, s.config.GrpcServer.Port))
	if err != nil {
		log.Fatalf("Failed to create listener: %v", err)
	}
	log.Printf("gRPC server listening on %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("Failed to serve grpc server : %v", err)
	}

}

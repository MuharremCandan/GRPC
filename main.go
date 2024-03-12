package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"test-grpc-project/pkg/config"
	"test-grpc-project/pkg/db"
	"test-grpc-project/pkg/gapi"
	"test-grpc-project/pkg/handler"
	"test-grpc-project/pkg/pb"
	"test-grpc-project/pkg/repository"
	"test-grpc-project/pkg/router"
	"test-grpc-project/pkg/service"

	"github.com/gofiber/fiber/v2"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"gorm.io/gorm"
)

func main() {

	config, err := config.LoadConfig("./")
	if err != nil {
		log.Fatalf("error loading config file %s:", err)
	}

	db := db.ConnectDb(&config)

	go runGatewayServer(db, config)
	gapi.NewServer(&config, db).NewgRpcServer()

}

func runGatewayServer(db *gorm.DB, config config.Config) {
	server := gapi.NewServer(&config, db)

	grpcMux := runtime.NewServeMux()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err := pb.RegisterGrpcProjectHandlerServer(ctx, grpcMux, server)

	if err != nil {
		log.Fatalf("failed to register handler server: %v", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	listener, err := net.Listen("tcp", net.JoinHostPort(config.HttpServer.Host, config.HttpServer.Port))
	if err != nil {
		log.Fatalf("failed to create a listener : %v", err)
	}
	log.Printf("start HTTP gateway server at %s", listener.Addr().String())
	err = http.Serve(listener, mux)

	if err != nil {
		log.Fatalf("failed to serve HTTP gateway server : %v", err)
	}
}

func runHttpServer(db *gorm.DB, config config.Config) {
	app := fiber.New()

	app.Get("/greet", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	userHandler := handler.NewUserHandler(service.NewUserService(repository.NewUserRepository(db)))

	router.NewRouter(userHandler).LoadRouter(app)

	app.Listen(net.JoinHostPort(config.HttpServer.Host, config.HttpServer.Port))
}

package main

import (
	u "TEMPLATE_MICROSERVICE/genproto/user"

	"TEMPLATE_MICROSERVICE/config"
	"TEMPLATE_MICROSERVICE/pkg/db"
	"TEMPLATE_MICROSERVICE/pkg/logger"
	"TEMPLATE_MICROSERVICE/service"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "golang")
	defer logger.Cleanup(log)

	connnDb, err := db.ConnectToDB(cfg)
	if err != nil {
		fmt.Println("Error while connecting postgres ", err.Error())
	}

	userService := service.NewUserService(connnDb, log)
	lis, err := net.Listen("tcp", cfg.UserServicePort)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	s := grpc.NewServer()
	reflection.Register(s)
	u.RegisterUserServiceServer(s, userService)

	log.Info("main: server is running",
		logger.String("port", cfg.UserServicePort))

	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

}

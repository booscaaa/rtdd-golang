package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/booscaaa/rtdd-golang/microservices/authenticator/adapter/grpcservice/accountservice"
	"github.com/booscaaa/rtdd-golang/microservices/authenticator/adapter/postgres"
	"github.com/booscaaa/rtdd-golang/microservices/authenticator/adapter/postgres/accountrepository"
	"github.com/booscaaa/rtdd-golang/microservices/authenticator/core/accountusecase"
	"github.com/booscaaa/rtdd-golang/microservices/authenticator/core/domain"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	ctx := context.Background()

	lis, err := net.Listen("tcp", ":3333")
	if err != nil {
		log.Fatal(err)
	}
	defer lis.Close()

	conn := postgres.GetConnection(ctx)
	defer conn.Close()

	accountRepository := accountrepository.NewAccountRepository(conn)
	accountUseCase := accountusecase.NewAccountUseCase(accountRepository)
	accountService := accountservice.NewAccountService(accountUseCase)
	grpcServer := grpc.NewServer()

	domain.RegisterAccountServiceServer(grpcServer, accountService)

	if err := grpcServer.Serve(lis); err != nil {
		fmt.Println(err)
	}
}

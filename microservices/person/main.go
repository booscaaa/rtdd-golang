package main

import (
	"fmt"
	"log"
	"net"

	"github.com/booscaaa/rtdd-golang/microservices/authenticator/adapter/grpcservice/personservice"
	"github.com/booscaaa/rtdd-golang/microservices/authenticator/adapter/mongodb"
	"github.com/booscaaa/rtdd-golang/microservices/authenticator/adapter/mongodb/personrepository"
	"github.com/booscaaa/rtdd-golang/microservices/authenticator/core/domain"
	"github.com/booscaaa/rtdd-golang/microservices/authenticator/core/personusecase"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":1111")
	if err != nil {
		log.Fatal(err)
	}
	defer lis.Close()

	db := mongodb.GetConnection()

	personRepository := personrepository.NewPersonRepository(db)
	personUseCase := personusecase.NewPersonUseCase(personRepository)
	personService := personservice.NewPersonService(personUseCase)
	grpcServer := grpc.NewServer()

	domain.RegisterPersonServiceServer(grpcServer, personService)

	if err := grpcServer.Serve(lis); err != nil {
		fmt.Println(err)
	}
}

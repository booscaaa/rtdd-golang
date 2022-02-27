package main

import (
	"fmt"
	"net"

	"github.com/booscaaa/rtdd-golang/microservices/person/adapter/grpcservice/personservice"
	"github.com/booscaaa/rtdd-golang/microservices/person/adapter/sqlite"
	"github.com/booscaaa/rtdd-golang/microservices/person/adapter/sqlite/personrepository"
	"github.com/booscaaa/rtdd-golang/microservices/person/core/domain"
	"github.com/booscaaa/rtdd-golang/microservices/person/core/personusecase"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":1111")
	if err != nil {
		fmt.Println(err)
	}
	defer lis.Close()

	db := sqlite.GetConnection()
	sqlite.ConfigConnection(db)

	personRepository := personrepository.NewPersonRepository(db)
	personUseCase := personusecase.NewPersonUseCase(personRepository)
	personService := personservice.NewPersonService(personUseCase)
	grpcServer := grpc.NewServer()

	domain.RegisterPersonServiceServer(grpcServer, personService)

	if err := grpcServer.Serve(lis); err != nil {
		fmt.Println(err)
	}
}

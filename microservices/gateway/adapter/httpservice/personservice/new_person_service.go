package personservice

import (
	"github.com/booscaaa/rtdd-golang/microservices/gateway/core/domain"
	"google.golang.org/grpc"
)

type service struct {
	grpc domain.PersonServiceClient
}

// NewPersonService .
func NewPersonService(conn *grpc.ClientConn) domain.PersonService {
	return &service{
		grpc: domain.NewPersonServiceClient(conn),
	}
}

package personservice

import (
	localDomain "github.com/booscaaa/rtdd-golang/microservices/gateway/core/domain"
	personDomain "github.com/booscaaa/rtdd-golang/microservices/person/core/domain"
	"google.golang.org/grpc"
)

type service struct {
	grpc personDomain.PersonServiceClient
}

// NewPersonService .
func NewPersonService(conn *grpc.ClientConn) localDomain.PersonService {
	return &service{
		grpc: personDomain.NewPersonServiceClient(conn),
	}
}

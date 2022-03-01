package accountservice

import (
	"github.com/booscaaa/rtdd-golang/microservices/gateway/core/domain"
	"google.golang.org/grpc"
)

type service struct {
	grpc domain.AccountServiceClient
}

// NewProductService .
func NewAccountService(conn *grpc.ClientConn) domain.AccountService {
	return &service{
		grpc: domain.NewAccountServiceClient(conn),
	}
}

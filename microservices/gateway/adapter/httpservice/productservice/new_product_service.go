package productservice

import (
	"github.com/booscaaa/rtdd-golang/microservices/gateway/core/domain"
	"google.golang.org/grpc"
)

type service struct {
	grpc domain.ProductServiceClient
}

// NewProductService .
func NewProductService(conn *grpc.ClientConn) domain.ProductService {
	return &service{
		grpc: domain.NewProductServiceClient(conn),
	}
}

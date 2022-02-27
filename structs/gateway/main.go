package main

import (
	"github.com/booscaaa/rtdd-golang/gateway/adapter/grpcservice"
	"github.com/booscaaa/rtdd-golang/gateway/adapter/httpservice"
)

func main() {
	go grpcservice.Run()
	httpservice.Run()
}

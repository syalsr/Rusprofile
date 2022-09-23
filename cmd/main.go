package main

import (
	_ "Rusprofile/docs"
	"Rusprofile/pkg"
	"Rusprofile/proto"
	"context"
	"log"
	"sync"
)

type RusprofileWrapper struct {
	proto.UnimplementedRusprofileWrapperServer
}

var (
	addr = "127.0.0.1:50051"
)

// @Summary Get
// @Description Get company info by INN
// @Router /INN [get]
// @Param input body proto.Request true "INN"
// @Success      200  {object}  proto.Response
func (receiver *RusprofileWrapper) Get(ctx context.Context, in *proto.Request) (
	*proto.Response,
	error,
) {
	response, err := pkg.FindCompanyByINN(in.GetINN())
	if err != nil {
		log.Print("No company wasnt found by this INN")

	}
	return response, nil
}

// @title Rusprofile
// @verion 1.0
// @description Get company info by INN

// @host localhost:8081
// @BasePath /
func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	go StartGrpcServer(&wg)
	go StartHppServe(&wg)
	go StartSwagger(&wg)

	wg.Wait()
}

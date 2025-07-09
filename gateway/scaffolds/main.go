package main

import (
	"event_sourcing_gateway/scaffolds/generator"
	"event_sourcing_gateway/scaffolds/model"
	"event_sourcing_gateway/scaffolds/parser"
	"event_sourcing_gateway/scaffolds/utils"
	"fmt"
)

const PROTO_PATH = "proto"

func main() {
	protoPaths := utils.GetAllProtoFiles(PROTO_PATH)

	protos := []*model.Proto{}
	for _, path := range protoPaths {
		proto := parser.ParseProtoFile(path)
		protos = append(protos, proto)
	}
	handlers, handlerPaths, err := generator.GenerateHandlers(protos)
	if err != nil {
		fmt.Println("failed to generate handlers:", err)
	}

	err = generator.GenerateRoutingRegistry(handlers, handlerPaths)
	if err != nil {
		fmt.Println("failed to generate routing registry:", err)
	}

	err = generator.GenerateGinRoutes(protos)
	if err != nil {
		fmt.Println("failed to generate gin routes:", err)
	}

	err = generator.GenerateGrpcClients(protos)
	if err != nil {
		fmt.Println("failed to generate grpc clients:", err)
	}

	err = generator.GenerateServiceClientRegistry(protos)
	if err != nil {
		fmt.Println("failed to generate service client registry:", err)
	}

	err = generator.GenerateInitRoutes(protos)
	if err != nil {
		fmt.Println("failed to generate init server routing:", err)
	}
}

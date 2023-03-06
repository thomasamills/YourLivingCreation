package grpc_server

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"testserver/conversation_code"
	humanize_protobuf "testserver/src/generated/humanize-protobuf"
)

func ServeMiddlewareAPI(address string, apiServer *conversation_code.MiddlewareApiImplementation) error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	//Db and Task List initialized here
	humanize_protobuf.RegisterConnectServer(grpcServer, apiServer)
	fmt.Println("Serving Middleware Grpc")
	if err := grpcServer.Serve(lis); err != nil {
		return err
	}
	return nil
}

func ServeCreationAPI(address string, apiServer *conversation_code.CreationApiImpl) error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	//Db and Task List initialized here
	humanize_protobuf.RegisterRuleApiServer(grpcServer, apiServer)
	fmt.Println("Serving Creation API Grpc")
	if err := grpcServer.Serve(lis); err != nil {
		return err
	}
	return nil
}

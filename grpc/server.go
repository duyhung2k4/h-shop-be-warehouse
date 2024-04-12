package grpc

import (
	"app/grpc/api"
	"app/grpc/proto"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func RunServerGRPC() {
	listenerGRPC, err := net.Listen("tcp", ":20005")

	if err != nil {
		log.Fatalln(listenerGRPC)
	}

	creds, errKey := credentials.NewServerTLSFromFile(
		"keys/server-warehouse/public.pem",
		"keys/server-warehouse/private.pem",
	)

	if errKey != nil {
		log.Fatalln(errKey)
	}

	grpcServer := grpc.NewServer(grpc.Creds(creds))
	proto.RegisterWarehouseServiceServer(grpcServer, api.NewWarehouseGRPC())

	log.Fatalln(grpcServer.Serve(listenerGRPC))
}

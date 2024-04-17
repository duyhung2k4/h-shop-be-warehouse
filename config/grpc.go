package config

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func connectGPRCServerProduct() {
	var errConn error

	creds, errKey := credentials.NewClientTLSFromFile("keys/server-product/public.pem", "localhost")
	if errKey != nil {
		log.Fatalln(errKey)
	}

	clientProduct, errConn = grpc.Dial(host+":20003", grpc.WithTransportCredentials(creds))
	if errConn != nil {
		log.Fatalln(errConn)
	}
}

package config

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func connectGPRCServerFile() {
	var errConn error

	creds, errKey := credentials.NewClientTLSFromFile("keys/server-file/public.pem", "localhost")
	if errKey != nil {
		log.Fatalln(errKey)
	}

	clientFile, errConn = grpc.Dial(host+":20004", grpc.WithTransportCredentials(creds))
	if errConn != nil {
		log.Fatalln(errConn)
	}
}

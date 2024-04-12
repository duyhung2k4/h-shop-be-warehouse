package config

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func connectGPRCServerFile() {
	var errProfile error

	creds, errKey := credentials.NewClientTLSFromFile("keys/server-file/public.pem", "localhost")
	if errKey != nil {
		log.Fatalln(errKey)
	}

	clientFile, errProfile = grpc.Dial(host+":20004", grpc.WithTransportCredentials(creds))
	if errProfile != nil {
		log.Fatalln(errProfile)
	}
}

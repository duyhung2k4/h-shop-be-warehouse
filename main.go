package main

import (
	"app/config"
	"app/router"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		server := http.Server{
			Addr:         ":" + config.GetAppPort(),
			Handler:      router.Router(),
			ReadTimeout:  time.Second * 30,
			WriteTimeout: time.Second * 30,
		}

		log.Fatalln(server.ListenAndServe())
		wg.Done()
	}()

	wg.Wait()
}

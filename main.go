package main

import (
	"example/pkg/settings"
	"example/routers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", settings.HTTPPort),
		Handler:        router,
		ReadTimeout:    settings.ReadTimeOut,
		WriteTimeout:   settings.WriteTimeOut,
		MaxHeaderBytes: 1 << 20,
	}

	err := s.ListenAndServe()
	if err != nil{
		log.Fatalf("Failed listen and provide service: %v", err)
	}
}

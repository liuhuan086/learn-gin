package main

import (
	"context"
	"example/pkg/settings"
	"example/routers"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	//router := routers.InitRouter()
	//
	//s := &http.Server{
	//	Addr:           fmt.Sprintf(":%d", settings.HTTPPort),
	//	Handler:        router,
	//	ReadTimeout:    settings.ReadTimeOut,
	//	WriteTimeout:   settings.WriteTimeOut,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//
	//err := s.ListenAndServe()
	//if err != nil{
	//	log.Fatalf("Failed listen and provide service: %v", err)
	//}

	/*
	endless 热更新是采取创建子进程后，将原进程退出的方式，这点不符合守护进程的要求
	 */
	//endless.DefaultReadTimeOut = settings.ReadTimeOut
	//endless.DefaultWriteTimeOut = settings.WriteTimeOut
	//endless.DefaultMaxHeaderBytes = 1 << 20
	//
	//endPoint := fmt.Sprintf(":%d", settings.HTTPPort)
	//
	//server := endless.NewServer(endPoint, routers.InitRouter())
	//server.BeforeBegin = func(add string) {
	//	log.Printf("Actual pid is %d", syscall.Getpid())
	//}
	//
	//err := server.ListenAndServe()
	//if err != nil {
	//	log.Printf("Server err: %v", err)
	//}

	/*
	http.Server - Shutdown()
	如果你的Golang >= 1.8，也可以考虑使用 http.Server 的 Shutdown 方法
	 */
	router := routers.InitRouter()
	
	s := &http.Server{
		Addr:              fmt.Sprintf("%d", settings.HTTPPort),
		Handler:           router,
		ReadTimeout:       settings.ReadTimeOut,
		WriteTimeout:      settings.WriteTimeOut,
		MaxHeaderBytes:    1 << 20,
	}

	go func() {
		if err := s.ListenAndServe();err != nil{
			log.Printf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<- quit

	log.Println("Shutdown Server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil{
		log.Fatal("Server shutdown: ", err)
	}

	log.Println("Server exiting")
}

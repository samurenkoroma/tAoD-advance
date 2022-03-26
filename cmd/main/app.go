package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net"
	"net/http"
	"tAoD-advance/internal/user"
	"time"
)

func main() {
	log.Printf("create router")
	router := httprouter.New()
	user.NewHandler().Register(router)
	start(router)

}

func start(router *httprouter.Router) {
	listener, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		panic(err)
	}
	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
	}
	log.Printf("run server")
	log.Fatal(server.Serve(listener))
}

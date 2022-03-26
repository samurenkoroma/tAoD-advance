package main

import (
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
	"tAoD-advance/internal/user"
	"tAoD-advance/pkg/logging"
	"time"
)

func main() {
	logger := logging.GetLogger()
	logger.Tracef("create router")
	router := httprouter.New()
	user.NewHandler(logger).Register(router)
	start(router)

}

func start(router *httprouter.Router) {
	logger := logging.GetLogger()
	listener, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		panic(err)
	}
	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
	}
	logger.Printf("run server")
	logger.Fatal(server.Serve(listener))
}

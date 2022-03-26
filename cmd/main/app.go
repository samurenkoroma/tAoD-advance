package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"tAoD-advance/internal/config"
	"tAoD-advance/internal/user"
	"tAoD-advance/pkg/logging"
	"time"
)

func main() {
	logger := logging.GetLogger()
	logger.Tracef("create router")
	router := httprouter.New()

	cfg := config.GetConfig()
	user.NewHandler(logger).Register(router)
	start(router, cfg)

}

func start(router *httprouter.Router, cfg *config.Config) {
	logger := logging.GetLogger()
	var listener net.Listener
	var listenError error
	if cfg.Listen.Type == "sock" {
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			logger.Fatal(err)
		}
		logger.Info("create socket")
		socketPath := path.Join(appDir, "app.sock")
		logger.Debugf("socket path: %s", socketPath)
		logger.Info("listen unix socket")
		listener, listenError = net.Listen("unix", socketPath)
		logger.Printf("run server on %s", socketPath)
	} else {
		logger.Info("listen tcp")
		listener, listenError = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIp, cfg.Listen.Port))
		logger.Printf("run server %s:%s", cfg.Listen.BindIp, cfg.Listen.Port)

	}
	if listenError != nil {
		logger.Fatal(listenError)
	}
	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
	}
	logger.Fatal(server.Serve(listener))
}

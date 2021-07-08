package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"petProject/internal/config"
	"petProject/internal/routes"
	"syscall"
	"time"
)

// @title PetProject API
// @version 1.0
// @description REST API for PetProject App

// @host localhost:8080
// @BasePath /api/v1/

// Run initializes whole application.
func Run() {
	cfg := config.Init()

	log.SetFormatter(&log.TextFormatter{ForceColors: true, ForceQuote: true})
	if cfg.ServerMode == config.Prod {
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(log.DebugLevel)
	}

	handler := routes.NewHandler()

	go func(srv *gin.Engine) {
		err := srv.Run(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port))
		if err != nil {
			log.Fatal(err.Error())
		}
	}(handler.Init(cfg))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Warn("Shutting down server...")
	time.Sleep(time.Second * 5)
}

package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"rhyme-engine/internal"
	"rhyme-engine/internal/config"
	"rhyme-engine/internal/logger"
	"rhyme-engine/internal/rhyme"
	"rhyme-engine/internal/rhyme/middleware"
	"rhyme-engine/internal/rhyme/service"
	"rhyme-engine/internal/wrapper"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

const (
	defaultCorrelationId = "00000000.00000000"
)

var ctx context.Context

func init() {
	ctx = internal.SetContextWithValue(context.Background(), internal.ContextKeyCorrelationID, defaultCorrelationId)
}

func main() {
	if os.Getenv(config.Env) == "" {
		configErr := config.ReadConfig(viper.GetString(config.Env))
		if configErr != nil {
			logger.Fatalf(ctx, "error in building configuration: %v", configErr)
		}
	}

	server := &http.Server{
		Addr:    ":" + viper.GetString(config.Port),
		Handler: createRouter(),
	}
	start(server)
}

func createRouter() *mux.Router {
	r := mux.NewRouter()
	r.Use(middleware.SetContentTypeHeader)
	reader := wrapper.FileReader{}
	refRhymeWords := service.NewReferenceRhymeSvcImpl(ctx, reader)
	baseRouter := r.PathPrefix(viper.GetString(config.BaseURL)).Subrouter()
	rhyme.Configure(baseRouter, refRhymeWords)
	return r
}

func start(server *http.Server) {
	go func() {
		logger.Infof(ctx, "Starting server on Port: %v\n", server.Addr)
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	}()
	gracefulStop(server)
}

func gracefulStop(server *http.Server) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	<-stop

	logger.Infof(ctx, "Shutting the server down...")

	ctx, cancelFunc := internal.NewContextWithTimeOut(ctx, 60*time.Second)
	defer cancelFunc()
	if err := server.Shutdown(ctx); err != nil {
		logger.Infof(ctx, "Error: %v\n", err)
	} else {
		logger.Infof(ctx, "Server stopped")
	}
}

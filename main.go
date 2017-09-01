package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/julienschmidt/httprouter"
	"github.com/laincloud/todomvc/api"
	"github.com/laincloud/todomvc/config"
	"go.uber.org/zap"
)

var (
	configFile = flag.String("config", "", "The configuration file")
)

func init() {
	flag.Parse()
}

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("zap.NewProduction() failed, error: %v.", err)
	}
	defer logger.Sync()

	if *configFile == "" {
		logger.Fatal("*configFile == \"\".")
	}

	c, err := config.New(*configFile)
	if err != nil {
		logger.Fatal("config.New() failed.",
			zap.String("ConfigFile", *configFile),
			zap.Error(err),
		)
	}

	db, err := xorm.NewEngine("mysql", c.MySQL.DataSourceName())
	if err != nil {
		logger.Fatal("xorm.NewEngine() failed.",
			zap.Error(err),
		)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM)

	router := httprouter.New()
	router.GET("/ping", api.Handle(api.Ping, db, logger))
	router.GET("/todos", api.Handle(api.GetTodos, db, logger))
	router.POST("/todos", api.Handle(api.CreateTodo, db, logger))
	router.GET("/todos/:id", api.Handle(api.GetTodo, db, logger))
	router.PATCH("/todos/:id", api.Handle(api.UpdateTodo, db, logger))
	router.DELETE("/todos/:id", api.Handle(api.DeleteTodo, db, logger))
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	go func() {
		if err1 := server.ListenAndServe(); err1 != nil {
			logger.Error("server.ListenAndServe() failed",
				zap.String("Addr", ":8080"),
				zap.Error(err),
			)
		}
	}()
	ctx := context.Background()
	defer func() {
		if err1 := server.Shutdown(ctx); err1 != nil {
			logger.Error("server.Shutdown() failed.",
				zap.String("Addr", ":8080"),
				zap.Error(err),
			)
		}
	}()

	<-quit
	logger.Info("Shutting down...")
}

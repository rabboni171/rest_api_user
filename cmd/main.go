package main

import (
	"context"
	"restApi/configs"
	"restApi/internal/handler"
	"restApi/internal/repository"
	"restApi/internal/service"
	"restApi/server"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	cfg, err := configs.InitConfig()
	if err != nil {
		logrus.Fatalf("error while initializing configs: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(cfg)
	if err != nil {
		logrus.Fatalf("Failed to initialize db: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewUserService(*repos)
	handlers := handler.NewUserHandler(*services)
	srv := new(server.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server %s", err.Error())
	}
	logrus.Print("user_service Started...")

	logrus.Print("user_service Shutting Down... Zzzzz")
	err = srv.Shutdown(context.Background())
	if err != nil {
		logrus.Errorf("error occured on server shutting down %s", err.Error())
	}
}

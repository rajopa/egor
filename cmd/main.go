package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	watcher "github.com/egor/watcher"
	"github.com/egor/watcher/pkg/handler"
	"github.com/egor/watcher/pkg/repository"
	"github.com/egor/watcher/pkg/service"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	slog.SetDefault(logger)

	if err := initConfig(); err != nil {
		logger.Error("error initializing configs: ", "err", err.Error())
		os.Exit(1)
	}
	if err := godotenv.Load(); err != nil {
		logger.Error("error loading env variables: ", "err", err.Error())
		os.Exit(1)
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBname:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logger.Error("failed to initialize db", "err", err.Error())
		os.Exit(1)
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos, logger)
	handlers := handler.NewHandler(services)

	srv := new(watcher.Server)
	go func() {
		services.Worker.Start()
	}()
	logger.Info("DomainApp Started")

	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logger.Error("error occurred while running http server:", "err", err.Error())
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logger.Info("DomainApp Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logger.Error("error occured on db connection close: ", "err", err.Error())
	}
	if err := db.Close(); err != nil {
		logger.Error("error occured on db connection close: ", "err", err.Error())
	}
}
func initConfig() error {
	viper.AddConfigPath("pkg/config")
	viper.SetConfigName("configs")
	return viper.ReadInConfig()
}

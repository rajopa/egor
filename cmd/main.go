package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	watcher "github.com/egor/watcher"
	"github.com/egor/watcher/kafka"
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
		Host:     viper.GetString("DB_HOST"),
		Port:     viper.GetString("DB_PORT"),
		Username: viper.GetString("DB_USER"),
		DBname:   viper.GetString("DB_NAME"),
		SSLMode:  viper.GetString("DB_SSLMODE"),
		Password: viper.GetString("DB_PASSWORD"),
	})
	if err != nil {
		logger.Error("failed to initialize db", "err", err.Error())
		os.Exit(1)
	}

	logger.Info("Attempting to apply migrations...")

	schema, err := os.ReadFile("schema/000001_init.up.sql")
	if err != nil {
		logger.Error("CRITICAL: migration file not found", "path", "schema/000001_init.up.sql", "err", err)
	} else {
		_, err = db.Exec(string(schema))
		if err != nil {
			logger.Error("failed to apply migration", "err", err)
		} else {
			logger.Info("Migrations applied successfully!")
		}
	}
	logger.Info("Waiting for DB to settle down...")
	time.Sleep(5 * time.Second)
	kafkaAddr := viper.GetString("KAFKA_BROKERS")
	if kafkaAddr == "" {
		kafkaAddr = "localhost:9092"
	}
	kafkaProducer := kafka.NewProducer([]string{kafkaAddr}, "target_updates", logger)
	repos := repository.NewRepository(db)
	services := service.NewService(repos, logger, kafkaProducer)
	handlers := handler.NewHandler(services)

	srv := new(watcher.Server)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go services.Worker.Start(ctx)

	go func() {
		logger.Info("DomainApp Started")

		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logger.Error("error occurred while running http server:", "err", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	cancel()

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
	viper.AutomaticEnv()
	return viper.ReadInConfig()
}

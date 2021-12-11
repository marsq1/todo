package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"os"
	"todo"
	"todo/handler"
	"todo/repository"
	"todo/service"
)

func main() {
	if err := initConfigs(); err != nil {
		log.Fatalf("Error init configs: %s", err)
	}

	if err := godotenv.Load("../.env"); err !=nil {
		log.Fatalf("Error load ..env file: %s", err)
	}

	db,err := repository.NewPostgresDB(repository.Config{
		Port: viper.GetString("db.port"),
		Host: viper.GetString("db.host"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName: viper.GetString("db.dbname"),
		SSLMode: viper.GetString("db.sslmode"),
	})

	if err != nil {
		log.Fatalf("Error database connect: %s", err)
	}

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)

	if err := srv.Run(viper.GetString("port"), handlers.InitRouters()); err != nil {
		log.Fatalf("Error run server: %s", err)
	}
}

func initConfigs() error{
	viper.AddConfigPath("../configs")
	viper.SetConfigName("../configs/config")
	return viper.ReadInConfig()
}

package main

import (
	"WEB_REST_exm0302"
	"WEB_REST_exm0302/pkg/cash"
	"WEB_REST_exm0302/pkg/handler"
	"WEB_REST_exm0302/pkg/mynats"
	"WEB_REST_exm0302/pkg/repository"
	"WEB_REST_exm0302/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {

		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	testCash := cash.NewCashTest()
	repos := repository.NewRepository(db)
	services := service.NewService(repos, testCash)
	handlers := handler.NewHandler(services)
	subsNats := mynats.NewSubsNats(services)

	errRecovery := services.RecoverCash()
	if errRecovery != nil {
		logrus.Fatalf("Ошибка восстановления кеша: %s", errRecovery)
	}
	srv := new(WEB_REST_exm0302.Server)

	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes(), subsNats); err != nil {
		logrus.Fatalf("error occured while running http server")
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}

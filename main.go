package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/spf13/viper"

	"github.com/betawulan/efishery/delivery"
	"github.com/betawulan/efishery/repository"
	"github.com/betawulan/efishery/service"
)

func main() {
	viper.AutomaticEnv()
	viper.SetConfigType("env")
	viper.SetConfigName(".env")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("failed running because file .env")
	}

	dsn := viper.GetString("mysql_dsn")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("can't connect database")
	}

	secretKey := viper.GetString("secret_key")

	authRepo := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepo, []byte(secretKey))

	e := echo.New()
	delivery.AddAuthRoute(authService, e)

	e.Logger.Fatal(e.Start(":5050"))
}

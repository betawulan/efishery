package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/spf13/viper"

	"github.com/betawulan/efishery/delivery"
	"github.com/betawulan/efishery/repository"
	"github.com/betawulan/efishery/service"
)

//	@title			Swagger Documentation APIs
//	@description	Documentation for Auth App and Fetch App
//	@host			localhost:5050
//	@BasePath		/auth
//	@Accept			json
//	@Produce		json
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
	urlCurrency := viper.GetString("url_currency")
	urlResource := viper.GetString("url_resource")
	port := viper.GetString("port")

	authRepo := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepo, []byte(secretKey))

	fishRepo := repository.NewFishRepository(urlCurrency, urlResource)
	fishService := service.NewFishService(fishRepo, urlCurrency, []byte(secretKey))

	e := echo.New()
	delivery.AddAuthRoute(authService, e)
	delivery.AddFishRoute(fishService, e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}

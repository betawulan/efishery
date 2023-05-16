package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/spf13/viper"

	"github.com/betawulan/efishery/delivery"
	"github.com/betawulan/efishery/packages/auth"
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

	registerRepo := repository.NewRegisterRepository(db)
	registerService := service.NewRegisterService(registerRepo)

	jwt := auth.NewAuth([]byte(secretKey))

	authRepo := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepo, jwt)

	e := echo.New()
	delivery.AddRegisterRoute(registerService, e)
	delivery.AddAuthRoute(authService, e)

	e.Logger.Fatal(e.Start(":5050"))
}

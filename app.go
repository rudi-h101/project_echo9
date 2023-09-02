package main

import (
	config "prakerja9/configs"
	"prakerja9/routes"
	"prakerja9/utils"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	LoadEnv()

	utils.InitDatabase()
	
	e := echo.New()
	
	routes.InitRoute(e)

	e.Logger.Fatal(e.Start(config.GetPort()))

}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}



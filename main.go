package main

import (
	config "r101/configs"
	"r101/routes"
	"r101/utils"

	"github.com/labstack/echo/v4"
)

func main() {
	// config.LoadEnv()

	utils.InitDatabase()
	
	e := echo.New()
	
	routes.InitRoute(e)

	e.Start(config.GetPort())

}



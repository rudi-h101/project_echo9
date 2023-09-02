package main

import (
	config "prakerja9/configs"
	"prakerja9/routes"
	"prakerja9/utils"

	"github.com/labstack/echo/v4"
)

func main() {
	config.LoadEnv()

	utils.InitDatabase()
	
	e := echo.New()
	
	routes.InitRoute(e)

	e.Start(config.GetPort())

}



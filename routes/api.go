package routes

import (
	"net/http"
	"os"
	"r101/controllers"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	echoSwagger "github.com/swaggo/echo-swagger"
)


func InitRoute(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RemoveTrailingSlash())

	e.GET("/", func(e echo.Context) error {
		return e.Redirect(http.StatusMovedPermanently, "/api-docs/index.html")
	})
	
	e.GET("/api-docs/*", echoSwagger.WrapHandler)
	
	e.POST("/register", controllers.Register)
	e.POST("/login", controllers.Login)

	// Restricted group
	r := e.Group("")
	// Configure middleware with the custom claims type
	r.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET"))))

	r.GET("/user-wishes", controllers.ListWish)
	r.POST("/user-wishes", controllers.CreateWish)
	r.GET("/user-wishes/:id", controllers.GetWish)
	r.PUT("/user-wishes/:id", controllers.UpdateWish)
	r.DELETE("/user-wishes/:id", controllers.DeleteWish)
	r.PUT("/user-wishes/:id/status", controllers.UpdateStatusWish)
}
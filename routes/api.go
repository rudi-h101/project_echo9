package routes

import (
	"os"
	"prakerja9/controllers"
	"prakerja9/models/dto"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


func InitRoute(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/register", controllers.Register)
	e.POST("/login", controllers.Login)

	// Restricted group
	r := e.Group("")
	// Configure middleware with the custom claims type
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(dto.JwtCustomClaims)
		},
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}
	r.Use(echojwt.WithConfig(config))

	r.GET("/user-wishes", controllers.ListWish)
	r.POST("/user-wishes", controllers.CreateWish)
	r.GET("/user-wishes/:id", controllers.GetWish)
	r.PUT("/user-wishes/:id", controllers.UpdateWish)
	r.DELETE("/user-wishes/:id", controllers.DeleteWish)
	r.PUT("/user-wishes/:id/status", controllers.UpdateStatusWish)
}
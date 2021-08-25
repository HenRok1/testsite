package api

import (
	"main/api/handlers"

	"github.com/labstack/echo"
)

func MainGroups(e *echo.Echo){
	e.GET("/hello", handlers.Hello)
	e.GET("/cats/:data", handlers.GetCats)
	e.GET("/login", handlers.Login)

	e.POST("/cats", handlers.AddCats)
	e.POST("/dogs", handlers.AddDog)
	e.POST("/hamster", handlers.AddHamster)
}

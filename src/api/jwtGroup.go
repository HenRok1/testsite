package api

import (
	"github.com/labstack/echo"
	"main/api/handlers"
)

func JWTGroup(g *echo.Group) {
	g.GET("/main", handlers.MainJwt)

}

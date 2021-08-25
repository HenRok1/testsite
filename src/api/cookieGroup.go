package api

import (
	"github.com/labstack/echo"
	"main/api/handlers"
)

func CookieGroup(g *echo.Group) {
	g.GET("/main", handlers.MainCookie)

}

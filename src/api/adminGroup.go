package api

import (
	"github.com/labstack/echo"
	"main/api/handlers"
)

func AdminGroup(g *echo.Group) {
	g.GET("/main", handlers.MainAdmin)

}

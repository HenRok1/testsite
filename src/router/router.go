package router

import (
	"github.com/labstack/echo"
	"main/api"
	"main/api/middlewares"
)

func New() *echo.Echo{
	e := echo.New()
	//create groups
	adminGroup:= e.Group("/admin")
	cookieGroup := e.Group("/cookie")
	jwtGroup := e.Group("/jwt")

	//set all middlewares
	middlewares.SetMainMiddlewares(e)
	middlewares.SetAdminMiddlewares(*adminGroup)
	middlewares.SetAdminMiddlewares(*cookieGroup)
	middlewares.SetAdminMiddlewares(*jwtGroup)

	//set main routes
	api.MainGroups(e)

	//set group routes

	api.AdminGroup(adminGroup)
	api.CookieGroup(cookieGroup)
	api.JWTGroup(jwtGroup)



	return e
}
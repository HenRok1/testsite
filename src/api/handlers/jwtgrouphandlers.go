package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

func MainJwt(c echo.Context) error{
	user := c.Get("user")

	token := user.(*jwt.Token)

	claims := token.Claims.(jwt.MapClaims)

	log.Println("User name: ", claims["name"], "User ID:", claims["jti"])

	return c.String(http.StatusOK, "you are on the top secret jwt page!")
}
package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"time"
)

type JwtClaims struct{
	Name 	string	`json:"name"`
	jwt.StandardClaims
}

func Login(c echo.Context) error{
	username := c.QueryParam("username")
	password := c.QueryParam("password")

	if username == "Nail" && password == "1234"{
		cookie := &http.Cookie{}
		//this is same
		//cookie:= 	new(http.Cookie)

		cookie.Name = "sessionID"
		cookie.Value = "some_string"
		cookie.Expires = time.Now().Add(48 * time.Hour)

		c.SetCookie(cookie)
		//Todo: create Jwt token
		token, err := createJwtToken()
		if err  != nil{
			log.Println("Error Creating JWT teken", err)
			return c.String(http.StatusInternalServerError, "something went wrong")
		}

		jwtCookie := &http.Cookie{}

		jwtCookie.Name = "JWTCookie"
		jwtCookie.Value = token
		jwtCookie.Expires = time.Now().Add(48 * time.Hour)

		c.SetCookie(jwtCookie)

		return c.JSON(http.StatusOK, map[string]string{
			"message": "You were logged in!",
			"token": token,
		})
	}
	return c.String(http.StatusUnauthorized, "Your username or password were wrong!")
}
func createJwtToken() (string, error){
	claims := JwtClaims{
		"Nail",
		jwt.StandardClaims{
			Id: "main_user_id",
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	token, err := rawToken.SignedString([]byte("mySecret"))
	if err != nil{
		return "", err
	}
	return token, err
}
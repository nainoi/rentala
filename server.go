package main

import (
	"net/http"
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"bitbucket.org/suthisakch/rentala/route"
	"github.com/dgrijalva/jwt-go"
)

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	log.Println(user)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	log.Println(name)
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

func main() {
	e := echo.New()

	route.Router(e)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("", restricted)

	e.Logger.Fatal(e.Start(":1323"))
}

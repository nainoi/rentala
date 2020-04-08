package middleware

import (
	"net/http"
	"time"

	response "bitbucket.org/suthisakch/rentala/api/v1"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	//"github.com/labstack/echo/middleware"
)

// IsAdmin as admin role
func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		isAdmin := claims["admin"].(bool)
		if isAdmin == false {
			//res.Response(http.StatusUnauthorized, "Unauthorized", nil)
			return echo.ErrUnauthorized
		}
		//res.Response(http.StatusOK, "", nil)
		return next(c)
	}
}

// IsMember as member role
func IsMember(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			res = response.Echo{C: c}
		)
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		role := claims["role"].(bool)
		if !role {
			res.Response(http.StatusUnauthorized, "Unauthorized", nil)
			return echo.ErrUnauthorized
		}
		res.Response(http.StatusOK, "", nil)
		return next(c)
	}
}

// jwtCustomClaims are custom claims extending default ones.
type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Throws unauthorized error
	if username != "jon" || password != "shhh!" {
		return echo.ErrUnauthorized
	}

	// Set custom claims
	claims := &jwtCustomClaims{
		"Jon Snow",
		true,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
	return err
	// if err != nil {
	// 	return err
	// }
}

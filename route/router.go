package route

import (
	"fmt"
	"net/http"

	"bitbucket.org/suthisakch/rentala/constants"

	auth "bitbucket.org/suthisakch/rentala/middleware"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	//response "bitbucket.org/suthisakch/rentala/api/v1"
)

//Router set router
func Router(e *echo.Echo) {

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	v1 := e.Group("/api/v1")
	adminGroup(v1)
	memberGroup(v1)
	v1.POST("/login", auth.Login)
	// e.Group("admin", auth.IsAdmin)
	// e.Group("member", auth.IsMember)
}

func adminGroup(g *echo.Group) {
	group := g.Group("/admin")
	group.Use(middleware.JWT([]byte("secret")), auth.IsAdmin)
	group.GET("", handler)
	fmt.Println(constants.ADMIN)
}

func memberGroup(g *echo.Group) {
	group := g.Group("/member", auth.IsMember)
	group.GET("", handler)
}

func handler(c echo.Context) error {
	// var (
	// 	res = response.Echo{C:c}
	// )
	//return res.Response(http.StatusUnauthorized,)
	return c.String(http.StatusOK, c.Request().RequestURI)
}

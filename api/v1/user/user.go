package user

import (
	token "bitbucket.org/suthisakch/rentala/middleware"
	"net/http"

	response "bitbucket.org/suthisakch/rentala/api/v1"
	"bitbucket.org/suthisakch/rentala/model"
	"bitbucket.org/suthisakch/rentala/repository"
	"github.com/labstack/echo"
)

type Handler struct {
	UserRepository repository.UserRepository
}

func (h Handler) CreateUser(c echo.Context) error {
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
}

func (h Handler) login(c echo.Context) {
	var (
		res = response.Echo{C: c}
	)
	var user model.User
	var err error

	var userEmail model.User
	err = c.Bind(userEmail)
	if err != nil {
		//context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		res.Response(http.StatusBadRequest, "user not found", nil)
	}
	user, err = h.UserRepository.GetUser(userEmail.Email, userEmail.Password)
	if err != nil {
		res.Response(http.StatusBadRequest, "user not found", nil)
	}
	tokenString, err := token.GenerateToken(user)
	if err != nil {
		res.Response(http.StatusBadRequest, "user not found", nil)
	}
	res.Response(http.StatusOK, "", tokenString)
}

// func (h Handler) getUser(c echo.Context) error {
// 	email := c.Param("email")
// 	if user == nil {
// 		return echo.NewHTTPError(http.StatusNotFound, "user not found")
// 	}
// 	return c.JSON(http.StatusOK, user)
// }

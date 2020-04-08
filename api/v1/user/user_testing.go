package user

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"bitbucket.org/suthisakch/rentala/model"
	"bitbucket.org/suthisakch/rentala/constants"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	userModel = model.User{
		UserID: primitive.NewObjectID(),
		FirstName:"test",
		LastName:"user",
		Role: constants.MEMBER,
		Email:"test@text.com",
	}
)

// TestCreateUser for add new user


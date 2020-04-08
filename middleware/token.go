package middleware

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"bitbucket.org/suthisakch/rentala/model"
)

//GenerateToken login user
func GenerateToken(user model.User) ( model.Token , error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	// This is the information which frontend can use
	// The backend can also decode the token and get admin etc.
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.UserID.String
	claims["firstname"] = user.FirstName
	claims["lastname"] = user.LastName
	claims["role"] = user.Role
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	// Generate encoded token and send it as response.
	// The signing string should be secret (a generated UUID works too)
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return model.Token{
			AccessToken :  "",
			RefreshToken: "",
		}, err
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	claims["id"] = user.UserID.String
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	rt, err := refreshToken.SignedString([]byte("secret"))
	if err != nil {
		return model.Token{
			AccessToken :  t,
			RefreshToken: rt,
		}, nil
	}

	return model.Token{
		AccessToken :  t,
		RefreshToken: rt,
	}, nil
}
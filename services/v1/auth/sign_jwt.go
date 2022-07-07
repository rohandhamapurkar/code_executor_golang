package auth

import (
	"rohandhamapurkar/code-executor/core/config"
	"rohandhamapurkar/code-executor/core/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func CreateJwtToken(user *models.User) (string, error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = user.ID
	atClaims["name"] = user.Name
	atClaims["email"] = user.Email
	atClaims["exp"] = time.Now().Add(time.Hour * 8).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	// Sign and get the complete encoded token as a string using the secret
	return token.SignedString([]byte(config.JwtSecret))

}

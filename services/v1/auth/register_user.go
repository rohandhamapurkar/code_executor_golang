package auth

import (
	"errors"
	"log"

	"rohandhamapurkar/code-executor/core/constants"
	"rohandhamapurkar/code-executor/core/db"
	"rohandhamapurkar/code-executor/core/models"
	"rohandhamapurkar/code-executor/core/structs"
	crypto "rohandhamapurkar/code-executor/libs/crypto"
)

// inserts the user into database and then returns access token
func RegisterUser(body *structs.RegUserReqBody) (string, error) {

	// get hashed password
	hashedPassword, err := crypto.EncryptPasswordString(body.Password)
	if err != nil {
		log.Println(err)
		return "", errors.New(constants.ERROR_PWD_HASH_FAILED)
	}

	body.Password = hashedPassword

	// create user model struct
	user := &models.User{Email: body.Email, Password: body.Password, Name: body.Name}

	// insert row into users table
	result := db.Postgres.Create(user)

	if result.Error != nil {
		log.Println(err)
		return "", errors.New(constants.ERROR_USER_INSERTION_FAILED)
	}

	// create jwt token
	jwtString, err := CreateJwtToken(user)
	if err != nil {
		log.Println(err)
		return "", errors.New(constants.ERROR_JWT_TOKEN_CREATION)
	}

	// return jwt access token string
	return jwtString, nil

}

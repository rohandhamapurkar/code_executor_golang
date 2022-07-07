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

func LoginUser(body *structs.LoginReqBody) (string, error) {

	var user *models.User
	result := db.Postgres.Find(&user, &models.User{Email: body.Email})

	if result.RowsAffected == 0 || result.Error != nil {
		log.Println("result.RowsAffected", result.RowsAffected)
		log.Println("result.Error", result.Error)
		return "", errors.New("Password invalid / No such user")
	}

	log.Println(user.Password, body.Password)
	err := crypto.CompareHashAndPassword(user.Password, body.Password)

	if err != nil {
		log.Println(err)
		return "", errors.New("Password invalid")
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

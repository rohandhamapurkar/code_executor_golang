package crypto

import (
	"errors"
	"log"
	"rohandhamapurkar/code-executor/core/constants"

	"golang.org/x/crypto/bcrypt"
)

func EncryptPasswordString(password string) (string, error) {
	// convert password to byte array
	passwordInBytes := []byte(password)

	// Hashing the password with the default cost of 10
	hashedPasswordInBytes, err := bcrypt.GenerateFromPassword(passwordInBytes, bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return "", errors.New(constants.ERROR_PWD_HASH_FAILED)
	}

	return string(hashedPasswordInBytes), nil
}

func CompareHashAndPassword(hashedPassword string, password string) error {

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err
}

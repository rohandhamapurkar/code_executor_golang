package auth

import (
	"rohandhamapurkar/code-executor/core/structs"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(user *structs.UserRegReqBody) error {
	// convert password to byte array
	passwordInBytes := []byte(user.Password)

	// Hashing the password with the default cost of 10
	hashedPasswordInBytes, err := bcrypt.GenerateFromPassword(passwordInBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// convert hashpassword byte array back to string
	user.Password = string(hashedPasswordInBytes)

	return nil
}

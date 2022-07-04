package auth

import (
	awsService "rohandhamapurkar/code-executor/services/v1/aws"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(email *string, password *string) error {
	passwordInBytes := []byte(*password)

	// Hashing the password with the default cost of 10
	hashedPasswordInBytes, err := bcrypt.GenerateFromPassword(passwordInBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	*password = string(hashedPasswordInBytes)

	err = awsService.SignUpUser(email, password)
	if err != nil {
		return err
	}
	return nil
}

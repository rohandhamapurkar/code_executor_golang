package auth

import (
	"errors"
	"rohandhamapurkar/code-executor/core/config"
	"rohandhamapurkar/code-executor/core/structs"

	"github.com/golang-jwt/jwt/v4"
)

func DecodeAndVerifyJwtToken(tokenString string) (*structs.JWTClaims, error) {
	token, err := jwt.Parse(tokenString, func(*jwt.Token) (interface{}, error) {
		return config.AwsCognitoJwtCachedPublicKey, nil
	})

	if err != nil {
		return &structs.JWTClaims{}, err
	}

	err = token.Claims.Valid()
	if err != nil {
		return &structs.JWTClaims{}, err
	}

	claims := &structs.JWTClaims{
		AuthTime: token.Claims.(jwt.MapClaims)["auth_time"].(float64),
		ClientId: token.Claims.(jwt.MapClaims)["client_id"].(string),
		Exp:      token.Claims.(jwt.MapClaims)["exp"].(float64),
		Iat:      token.Claims.(jwt.MapClaims)["iat"].(float64),
		Iss:      token.Claims.(jwt.MapClaims)["iss"].(string),
		Jti:      token.Claims.(jwt.MapClaims)["jti"].(string),
		Scope:    token.Claims.(jwt.MapClaims)["scope"].(string),
		Sub:      token.Claims.(jwt.MapClaims)["sub"].(string),
		TokenUse: token.Claims.(jwt.MapClaims)["token_use"].(string),
		Username: token.Claims.(jwt.MapClaims)["username"].(string),
		Version:  token.Claims.(jwt.MapClaims)["version"].(float64),
	}

	if claims.ClientId != config.AwsCognitoClientId {
		return &structs.JWTClaims{}, errors.New("Invalid Client Id")
	}

	if claims.Iss != config.AwsCognitoIssuer {
		return &structs.JWTClaims{}, errors.New("Invalid Issuer")
	}
	if claims.TokenUse != "access" {
		return &structs.JWTClaims{}, errors.New("Invalid Token Use")
	}

	return claims, nil
}

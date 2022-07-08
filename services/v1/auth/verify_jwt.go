package auth

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"rohandhamapurkar/code-executor/core/config"
	"rohandhamapurkar/code-executor/core/structs"

	"github.com/golang-jwt/jwt/v4"
)

func convertKey(rawE, rawN string) *rsa.PublicKey {
	decodedE, err := base64.RawURLEncoding.DecodeString(rawE)
	if err != nil {
		panic(err)
	}
	if len(decodedE) < 4 {
		ndata := make([]byte, 4)
		copy(ndata[4-len(decodedE):], decodedE)
		decodedE = ndata
	}
	pubKey := &rsa.PublicKey{
		N: &big.Int{},
		E: int(binary.BigEndian.Uint32(decodedE[:])),
	}
	decodedN, err := base64.RawURLEncoding.DecodeString(rawN)
	if err != nil {
		panic(err)
	}
	pubKey.N.SetBytes(decodedN)
	return pubKey
}

func DecodeAndVerifyJwtToken(tokenString string) (*structs.JWTClaims, error) {
	token, err := jwt.Parse(tokenString, func(*jwt.Token) (interface{}, error) {
		key := convertKey(config.AwsCognitoJwks.Keys[1].E, config.AwsCognitoJwks.Keys[1].N)
		return key, nil
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

func CacheJWK() error {
	req, err := http.NewRequest("GET", config.AwsCognitoJwksUrl, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	jwk := new(structs.JWK)
	err = json.Unmarshal(body, jwk)
	if err != nil {
		return err
	}

	config.AwsCognitoJwks = jwk

	log.Println("Cached AWS Cognito jwks")
	return nil
}

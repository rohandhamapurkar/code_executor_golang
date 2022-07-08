package auth

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
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

func DecodeAndVerifyJwtToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(*jwt.Token) (interface{}, error) {
		key := convertKey(config.AwsCognitoJwks.Keys[1].E, config.AwsCognitoJwks.Keys[1].N)
		return key, nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
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

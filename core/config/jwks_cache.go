package config

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"io/ioutil"
	"math/big"
	"net/http"
	"rohandhamapurkar/code-executor/core/structs"
)

func cacheAWSCognitoJWK() (*rsa.PublicKey, error) {
	req, err := http.NewRequest("GET", AwsCognitoJwksUrl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	jwk := new(structs.JWK)
	err = json.Unmarshal(body, jwk)
	if err != nil {
		return nil, err
	}

	var pubKey *rsa.PublicKey
	pubKey, err = convertKey(jwk.Keys[1].E, jwk.Keys[1].N)
	if err != nil {
		return nil, err
	}

	return pubKey, nil
}

func convertKey(rawE, rawN string) (*rsa.PublicKey, error) {
	decodedE, err := base64.RawURLEncoding.DecodeString(rawE)
	if err != nil {
		return nil, err
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
		return nil, err
	}
	pubKey.N.SetBytes(decodedN)
	return pubKey, nil
}

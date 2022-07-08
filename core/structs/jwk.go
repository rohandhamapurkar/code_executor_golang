package structs

type JWK struct {
	Keys []struct {
		Alg string `json:"alg"`
		E   string `json:"e"`
		Kid string `json:"kid"`
		Kty string `json:"kty"`
		N   string `json:"n"`
	} `json:"keys"`
}

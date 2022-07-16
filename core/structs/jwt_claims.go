package structs

type JWTClaims struct {
	AuthTime float64
	ClientID string
	Exp      float64
	Iat      float64
	Iss      string
	Jti      string
	Scope    string
	Sub      string
	TokenUse string
	Username string
	Version  float64
}

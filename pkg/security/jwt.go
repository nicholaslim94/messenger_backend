package security

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	Expire30Mins = time.Now().Add(30 * time.Minute).Unix()
	Expire10Days = time.Now().Add(240 * time.Hour).Unix()
	Issuer       = "System"
)

//JwtClaims structure As shown in the package documentation example
type JwtClaims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

func GetSecret() []byte {
	//TODO: GET SECRET
	return []byte("TEMPORARYSECRETKEYUSEYMLFILEINSTEAD")
}

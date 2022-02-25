package jwt

import (
	"os"
	"time"

	j "github.com/golang-jwt/jwt"
)

var JWT_SECRET []byte

func InitJWTToken() {
	var jwtSecretVar string = os.Getenv("JWT_SECRET")
	if jwtSecretVar == "" {
		panic("JWT_SECRET not set")
	}
	JWT_SECRET = []byte(jwtSecretVar)
}

func CreateToken(userid uint64) (string, error) {
	var jwtMapClaims j.MapClaims = j.MapClaims{}
	jwtMapClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	jwtMapClaims["permissions"] = [5]int{1, 2, 3, 4, 5}

	var jwtWithClaims *j.Token = j.NewWithClaims(j.SigningMethodHS256, jwtMapClaims)
	token, err := jwtWithClaims.SignedString(JWT_SECRET)
	if err != nil {
		return "", err
	}
	return token, nil
}

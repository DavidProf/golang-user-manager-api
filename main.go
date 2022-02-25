package main

import (
	"github.com/Not-the-optimal-solution-NTOS/golang-user-manager-api/jwt"
	"github.com/Not-the-optimal-solution-NTOS/golang-user-manager-api/password"
)

var JWT_SECRET []byte

func main() {

	jwt.InitJWTToken()

	token, error := jwt.CreateToken(1)
	if error != nil {
		panic(error)
	}
	println(token)
	println(password.CreatePasswordHash("password"))
	println("Hello, world!")
}

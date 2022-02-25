package user

import (
	pwd "github.com/Not-the-optimal-solution-NTOS/golang-user-manager-api/password"
)

type User struct {
	name        string
	email       string
	hash        string
	permissions []string
}

func CreateUser(name string, email string, password string) *User {
	var hash string = pwd.CreatePasswordHash(password)
	return &User{name, email, hash, []string{}}
}

func LoginUser(email string, password string) bool {
	return true
}

func findUser(email string) *User {
}

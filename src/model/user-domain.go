package model

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/Emanuelvss13/meu-primeiro-crud-go/src/configuration/Error"
)

type userDomain struct {
	Email    string
	Password string
	Name     string
	Age      int8
}

func NewUserDomain(
	email, password, name string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		Name:     name,
		Age:      age,
		Email:    email,
		Password: password,
	}
}

func (u *userDomain) EncryptPassword() {

	hash := md5.New()
	defer hash.Reset()

	hash.Write([]byte(u.Password))

	u.Password = hex.EncodeToString(hash.Sum(nil))

}

type UserDomainInterface interface {
	CreateUser() *Error.RestError
	UpdateUser(string) *Error.RestError
	FindUser(string) (*userDomain, *Error.RestError)
	DeleteUser(string) *Error.RestError
}

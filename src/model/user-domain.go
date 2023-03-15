package model

import (
	"crypto/md5"
	"encoding/hex"
)

type UserDomainInterface interface {
	GetName() string
	GetAge() int8
	GetEmail() string
	GetPassword() string

	EncryptPassword()
}

type userDomain struct {
	email    string
	password string
	name     string
	age      int8
}

func (u *userDomain) GetName() string {
	return u.name
}

func (u *userDomain) GetAge() int8 {
	return u.age
}

func (u *userDomain) GetEmail() string {
	return u.email
}

func (u *userDomain) GetPassword() string {
	return u.password
}

func NewUserDomain(
	email, password, name string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		name:     name,
		age:      age,
		email:    email,
		password: password,
	}
}

func (u *userDomain) EncryptPassword() {

	hash := md5.New()
	defer hash.Reset()

	hash.Write([]byte(u.password))

	u.password = hex.EncodeToString(hash.Sum(nil))

}

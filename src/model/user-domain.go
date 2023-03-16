package model

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
)

type UserDomainInterface interface {
	GetName() string
	GetAge() int8
	GetEmail() string
	GetPassword() string

	ChangeID(string)

	GetJSONValue() (string, error)
	EncryptPassword()
}

type userDomain struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Age      int8   `json:"age"`
}

func (u *userDomain) GetJSONValue() (string, error) {
	json, err := json.Marshal(u)

	if err != nil {

		return "", err
	}

	return string(json), nil
}

func (u *userDomain) GetName() string {
	return u.Name
}

func (u *userDomain) GetAge() int8 {
	return u.Age
}

func (u *userDomain) GetEmail() string {
	return u.Email
}

func (u *userDomain) GetPassword() string {
	return u.Password
}

func NewUserDomain(
	email, password, name string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		ID:       "",
		Name:     name,
		Age:      age,
		Email:    email,
		Password: password,
	}
}

func (u *userDomain) ChangeID(id string) {
	u.ID = id
}

func (u *userDomain) EncryptPassword() {

	hash := md5.New()
	defer hash.Reset()

	hash.Write([]byte(u.Password))

	u.Password = hex.EncodeToString(hash.Sum(nil))

}

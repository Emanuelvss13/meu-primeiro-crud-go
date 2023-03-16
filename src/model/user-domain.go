package model

import (
	"encoding/json"
)

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

func (u *userDomain) ChangeID(id string) {
	u.ID = id
}

package service

import (
	"github.com/Emanuelvss13/meu-primeiro-crud-go/src/configuration/Error"
	"github.com/Emanuelvss13/meu-primeiro-crud-go/src/model"
)

func (u *userDomainService) FindUserByID(
	id string,
) (model.UserDomainInterface, *Error.RestError) {
	return u.userRepository.FindUserByID(id)
}

func (u *userDomainService) FindUserByEmail(
	email string,
) (model.UserDomainInterface, *Error.RestError) {
	return u.userRepository.FindUserByEmail(email)
}

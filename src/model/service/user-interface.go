package service

import (
	"github.com/Emanuelvss13/meu-primeiro-crud-go/src/configuration/Error"
	"github.com/Emanuelvss13/meu-primeiro-crud-go/src/model"
)

type userDomainService struct{}

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type UserDomainService interface {
	CreateUser(userDomain model.UserDomainInterface) *Error.RestError
	UpdateUser(string, userDomain *model.UserDomainInterface) *Error.RestError
	FindUser(string) (*model.UserDomainInterface, *Error.RestError)
	DeleteUser(string) *Error.RestError
}

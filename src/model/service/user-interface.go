package service

import (
	"github.com/Emanuelvss13/meu-primeiro-crud-go/src/configuration/Error"
	"github.com/Emanuelvss13/meu-primeiro-crud-go/src/model"
	"github.com/Emanuelvss13/meu-primeiro-crud-go/src/model/repository"
)

type userDomainService struct {
	userRepository repository.UserRepositoryInterface
}

func NewUserDomainService(repository repository.UserRepositoryInterface) UserDomainService {
	return &userDomainService{userRepository: repository}
}

type UserDomainService interface {
	CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *Error.RestError)
	UpdateUser(string, userDomain *model.UserDomainInterface) *Error.RestError
	FindUser(string) (*model.UserDomainInterface, *Error.RestError)
	DeleteUser(string) *Error.RestError
}

package service

import (
	"fmt"

	"github.com/Emanuelvss13/meu-primeiro-crud-go/src/configuration/Error"
	"github.com/Emanuelvss13/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/Emanuelvss13/meu-primeiro-crud-go/src/model"
	"go.uber.org/zap"
)

func (u *userDomainService) CreateUser(userDomain model.UserDomainInterface) *Error.RestError {

	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()

	fmt.Println(userDomain)

	return nil
}

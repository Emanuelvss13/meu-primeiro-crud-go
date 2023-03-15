package model

import (
	"fmt"

	"github.com/Emanuelvss13/meu-primeiro-crud-go/src/configuration/Error"
	"github.com/Emanuelvss13/meu-primeiro-crud-go/src/configuration/logger"
	"go.uber.org/zap"
)

func (u *userDomain) CreateUser() *Error.RestError {

	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	u.EncryptPassword()

	fmt.Println(u)

	return nil
}

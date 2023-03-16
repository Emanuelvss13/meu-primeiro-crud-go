package controller

import (
	"github.com/Emanuelvss13/meu-primeiro-crud-go/src/model/service"
	"github.com/gin-gonic/gin"
)

type UserControllerInterface interface {
	FindUserByID(c *gin.Context)
	FindUserByEmail(c *gin.Context)

	CreateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	UpdateUser(c *gin.Context)
}

type userController struct {
	service service.UserDomainService
}

func NewUserController(
	service service.UserDomainService,
) UserControllerInterface {
	return &userController{service: service}
}

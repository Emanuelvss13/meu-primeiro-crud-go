package controller

import (
	"github.com/Emanuelvss13/meu-primeiro-crud-go/src/configuration/Error"
	"github.com/gin-gonic/gin"
)

func (uc *userController) FindUserByID(c *gin.Context) {

	err := Error.NewInternalServerError("Error")

	c.JSON(err.Code, err)

}

func (uc *userController) FindUserByEmail(c *gin.Context) {

}

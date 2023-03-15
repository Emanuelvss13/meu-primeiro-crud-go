package controller

import (
	"github.com/Emanuelvss13/meu-primeiro-crud-go/src/configuration/Error"
	"github.com/gin-gonic/gin"
)

func FindUserByID(c *gin.Context) {

	err := Error.NewInternalServerError("Error")

	c.JSON(err.Code, err)

}

func FindUserByEmail(c *gin.Context) {

}

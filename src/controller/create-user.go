package controller

import (
	"fmt"
	"net/http"

	"github.com/Emanuelvss13/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/Emanuelvss13/meu-primeiro-crud-go/src/configuration/validation"
	"github.com/Emanuelvss13/meu-primeiro-crud-go/src/controller/model/request"
	"github.com/Emanuelvss13/meu-primeiro-crud-go/src/model"
	"github.com/Emanuelvss13/meu-primeiro-crud-go/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userController) CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		fmt.Printf("There are some incorrect fieds, error=%s\n", err.Error())
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	createdUser, err := uc.service.CreateUser(domain)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User creted successfully", zap.String("journey", "createUser"))

	c.JSON(http.StatusOK, view.ConvetDomainToResponse(
		createdUser,
	))

}

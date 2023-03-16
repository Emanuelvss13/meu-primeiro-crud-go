package controller

import (
	"net/http"
	"net/mail"

	"github.com/Emanuelvss13/meu-primeiro-crud-go/src/configuration/Error"
	"github.com/Emanuelvss13/meu-primeiro-crud-go/src/view"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (uc *userController) FindUserByID(c *gin.Context) {

	userID := c.Param("userId")

	if _, err := primitive.ObjectIDFromHex(userID); err != nil {

		errorMessage := Error.NewBadRequestError("Invalid user ID")

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByID(userID)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, view.ConvetDomainToResponse(userDomain))
}

func (uc *userController) FindUserByEmail(c *gin.Context) {

	userEmail := c.Param("userEmail")

	if _, err := mail.ParseAddress(userEmail); err != nil {

		errorMessage := Error.NewBadRequestError("Invalid user Email")

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByEmail(userEmail)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, view.ConvetDomainToResponse(userDomain))
}

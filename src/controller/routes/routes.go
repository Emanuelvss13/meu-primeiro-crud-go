package routes

import (
	"github.com/Emanuelvss13/meu-primeiro-crud-go/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {

	r.GET("/getUserById/:userId", controller.FindUserByID)
	r.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	r.POST("/createUser")
	r.PATCH("/updateUser/:userId")
	r.DELETE("/deleteUser/:userId")

}

package routes

import (
	"api/src/controller"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, useController controller.UserControllerInterface) {
	r.GET("/userid/:userId", useController.FindUserById)
	r.GET("/email/:userEmail", useController.FindUserByEmail)
	r.POST("/create", useController.CreateUser)
	r.PUT("/edit/:userId", useController.EditUserById)
	r.DELETE("/excluiruser/:userId", useController.DeleteById)

}
